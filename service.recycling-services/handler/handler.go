package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/edstell/lambda/libraries/errors"
	"github.com/edstell/lambda/service.recycling-services/notifier"
	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"github.com/edstell/lambda/service.recycling-services/services"
	"github.com/edstell/lambda/service.recycling-services/store"
	twilio "github.com/edstell/lambda/service.twilio/rpc"
)

type handler struct {
	store           store.Store
	client          *twilio.Client
	fetcher         services.Fetcher
	timeNow         func() time.Time
	propertyMessage func(string, recyclingservicesproto.Property) (notifier.Message, error)
}

func New(store store.Store, client *twilio.Client, timeNow func() time.Time) recyclingservicesproto.Handler {
	return &handler{
		store:  store,
		client: client,
		fetcher: services.WebScraper(&http.Client{
			Timeout: time.Second * 30,
		},
			services.ParseHTML,
			"https://recyclingservicesproto.bromley.gov.uk/property",
		),
		timeNow:         timeNow,
		propertyMessage: propertyMessageFunc(timeNow),
	}
}

func (h *handler) ReadProperty(ctx context.Context, body recyclingservicesproto.ReadPropertyRequest) (*recyclingservicesproto.ReadPropertyResponse, error) {
	property, err := h.store.ReadProperty(ctx, body.PropertyId)
	if err != nil {
		return nil, err
	}
	return &recyclingservicesproto.ReadPropertyResponse{
		Property: property,
	}, nil
}

func (h *handler) SyncProperty(ctx context.Context, body recyclingservicesproto.SyncPropertyRequest) (*recyclingservicesproto.SyncPropertyResponse, error) {
	services, err := h.fetcher.Fetch(ctx, body.PropertyId)
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, errors.NewKnown(http.StatusInternalServerError, "failed to fetch any services")
	}
	property, err := h.store.WriteProperty(ctx, body.PropertyId, services)
	if err != nil {
		return nil, err
	}
	return &recyclingservicesproto.SyncPropertyResponse{
		Property: property,
	}, nil
}

func (h *handler) NotifyProperty(ctx context.Context, body recyclingservicesproto.NotifyPropertyRequest) (*recyclingservicesproto.NotifyPropertyResponse, error) {
	_, err := h.store.ReadProperty(ctx, body.PropertyId)
	if err != nil {
		return nil, err
	}

	// message, err := h.propertyMessage(body.Message, *property)
	// if err != nil {
	// 	return nil, err
	// }

	// sms := notifier.SMS(h.client, body.PhoneNumber)
	// if err := sms.Notify(ctx, message); err != nil {
	// 	return nil, err
	// }

	return &recyclingservicesproto.NotifyPropertyResponse{}, nil
}

func propertyMessageFunc(timeNow func() time.Time) func(string, *recyclingservicesproto.Property) (notifier.Message, error) {
	servicesTomorrow := notifier.ServicesTomorrow(timeNow)
	servicesNextWeek := notifier.ServicesNextWeek(timeNow)
	describeProperty := notifier.DescribeProperty()
	return func(typ string, property *recyclingservicesproto.Property) (notifier.Message, error) {
		switch typ {
		case recyclingservicesproto.MessageServicesTomorrow:
			return servicesTomorrow(property), nil
		case recyclingservicesproto.MessageServicesNextWeek:
			return servicesNextWeek(property), nil
		case recyclingservicesproto.MessageDescribeProperty:
			return describeProperty(property), nil
		default:
			return nil, errors.BadRequest("bad param: message")
		}
	}
}
