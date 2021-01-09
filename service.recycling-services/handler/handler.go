package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/edstell/lambda/libraries/errors"
	"github.com/edstell/lambda/service.recycling-services/notifier"
	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
	"github.com/edstell/lambda/service.recycling-services/services"
	"github.com/edstell/lambda/service.recycling-services/store"
	twilio "github.com/edstell/lambda/service.twilio/rpc"
)

type handler struct {
	store           store.Store
	client          *twilio.Client
	fetcher         services.Fetcher
	timeNow         func() time.Time
	propertyMessage func(string, recyclingservices.Property) (notifier.Message, error)
}

func New(store store.Store, client *twilio.Client, timeNow func() time.Time) recyclingservices.Handler {
	return &handler{
		store:  store,
		client: client,
		fetcher: services.WebScraper(&http.Client{
			Timeout: time.Second * 30,
		},
			services.ParseHTML,
			"https://recyclingservices.bromley.gov.uk/property",
		),
		timeNow:         timeNow,
		propertyMessage: propertyMessageFunc(timeNow),
	}
}

func (h *handler) ReadProperty(ctx context.Context, body recyclingservices.ReadPropertyRequest) (*recyclingservices.ReadPropertyResponse, error) {
	property, err := h.store.ReadProperty(ctx, body.PropertyID)
	if err != nil {
		return nil, err
	}
	return &recyclingservices.ReadPropertyResponse{
		Property: *property,
	}, nil
}

func (h *handler) WriteProperty(ctx context.Context, body recyclingservices.WritePropertyRequest) (*recyclingservices.WritePropertyResponse, error) {
	property, err := h.store.WriteProperty(ctx, body.PropertyID, body.Services)
	if err != nil {
		return nil, err
	}
	return &recyclingservices.WritePropertyResponse{
		Property: *property,
	}, nil
}

func (h *handler) SyncProperty(ctx context.Context, body recyclingservices.SyncPropertyRequest) (*recyclingservices.SyncPropertyResponse, error) {
	services, err := h.fetcher.Fetch(ctx, body.PropertyID)
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, errors.NewKnown(http.StatusInternalServerError, "failed to fetch any services")
	}
	property, err := h.store.WriteProperty(ctx, body.PropertyID, services)
	if err != nil {
		return nil, err
	}
	return &recyclingservices.SyncPropertyResponse{
		Property: *property,
	}, nil
}

func (h *handler) NotifyProperty(ctx context.Context, body recyclingservices.NotifyPropertyRequest) (*recyclingservices.NotifyPropertyResponse, error) {
	property, err := h.store.ReadProperty(ctx, body.PropertyID)
	if err != nil {
		return nil, err
	}

	message, err := h.propertyMessage(body.Message, *property)
	if err != nil {
		return nil, err
	}

	sms := notifier.SMS(h.client, body.PhoneNumber)
	if err := sms.Notify(ctx, message); err != nil {
		return nil, err
	}

	return &recyclingservices.NotifyPropertyResponse{}, nil
}

func propertyMessageFunc(timeNow func() time.Time) func(string, recyclingservices.Property) (notifier.Message, error) {
	servicesTomorrow := notifier.ServicesTomorrow(timeNow)
	servicesThisWeek := notifier.ServicesThisWeek(timeNow)
	describeProperty := notifier.DescribeProperty()
	return func(typ string, property recyclingservices.Property) (notifier.Message, error) {
		switch typ {
		case recyclingservices.MessageServicesTomorrow:
			return servicesTomorrow(property), nil
		case recyclingservices.MessageServicesThisWeek:
			return servicesThisWeek(property), nil
		case recyclingservices.MessageDescribeProperty:
			return describeProperty(property), nil
		default:
			return nil, errors.BadRequest("bad param: message")
		}
	}
}
