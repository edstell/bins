package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/edstell/lambda/service.recycling-services/domain"
	"github.com/edstell/lambda/service.recycling-services/notifier"
	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"github.com/edstell/lambda/service.recycling-services/services"
	"github.com/edstell/lambda/service.recycling-services/store"
	twilioproto "github.com/edstell/lambda/service.twilio/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	store             store.Store
	fetcher           services.Fetcher
	timeNow           func() time.Time
	propertyMessage   func(string, domain.Property) (notifier.Message, error)
	notifierFromProto func(*recyclingservicesproto.Notifier) notifier.Notifier
}

func New(store store.Store, twilio *twilioproto.Client, timeNow func() time.Time) recyclingservicesproto.Handler {
	return &handler{
		store: store,
		fetcher: services.WebScraper(
			&http.Client{Timeout: time.Second * 30},
			services.ParseHTML,
			"https://recyclingservicesproto.bromley.gov.uk/property",
		),
		timeNow:           timeNow,
		propertyMessage:   propertyMessageFunc(timeNow),
		notifierFromProto: notifier.FromProtoFunc(twilio),
	}
}

func (h *handler) ReadProperty(ctx context.Context, body *recyclingservicesproto.ReadPropertyRequest) (*recyclingservicesproto.ReadPropertyResponse, error) {
	property, err := h.store.ReadProperty(ctx, body.PropertyId)
	if err != nil {
		return nil, err
	}
	return &recyclingservicesproto.ReadPropertyResponse{
		Property: property.ToProto(),
	}, nil
}

func (h *handler) SyncProperty(ctx context.Context, body *recyclingservicesproto.SyncPropertyRequest) (*recyclingservicesproto.SyncPropertyResponse, error) {
	services, err := h.fetcher.Fetch(ctx, body.PropertyId)
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, status.Error(codes.Internal, "failed to fetch any services")
	}
	property, err := h.store.WriteProperty(ctx, body.PropertyId, services)
	if err != nil {
		return nil, err
	}
	return &recyclingservicesproto.SyncPropertyResponse{
		Property: property.ToProto(),
	}, nil
}

func (h *handler) NotifyProperty(ctx context.Context, body *recyclingservicesproto.NotifyPropertyRequest) (*recyclingservicesproto.NotifyPropertyResponse, error) {
	property, err := h.store.ReadProperty(ctx, body.PropertyId)
	if err != nil {
		return nil, err
	}
	notifier := h.notifierFromProto(body.Notifier)
	message, err := h.propertyMessage(body.MessageName, *property)
	if err != nil {
		return nil, err
	}
	if err := notifier.Notify(ctx, message); err != nil {
		return nil, err
	}
	return &recyclingservicesproto.NotifyPropertyResponse{}, nil
}

func propertyMessageFunc(timeNow func() time.Time) func(string, domain.Property) (notifier.Message, error) {
	servicesTomorrow := notifier.ServicesTomorrow(timeNow)
	servicesNextWeek := notifier.ServicesNextWeek(timeNow)
	describeProperty := notifier.DescribeProperty()
	return func(typ string, property domain.Property) (notifier.Message, error) {
		switch typ {
		case recyclingservicesproto.MessageServicesTomorrow:
			return servicesTomorrow(property)
		case recyclingservicesproto.MessageServicesNextWeek:
			return servicesNextWeek(property)
		case recyclingservicesproto.MessageDescribeProperty:
			return describeProperty(property)
		default:
			return nil, status.Error(codes.InvalidArgument, "")
		}
	}
}
