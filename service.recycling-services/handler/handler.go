package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	notifierproto "github.com/edstell/lambda/service.notifier/proto"
	"github.com/edstell/lambda/service.recycling-services/domain"
	"github.com/edstell/lambda/service.recycling-services/message"
	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"github.com/edstell/lambda/service.recycling-services/services"
	"github.com/edstell/lambda/service.recycling-services/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	store           store.Store
	fetcher         services.Fetcher
	timeNow         func() time.Time
	notifier        notifierproto.Client
	propertyMessage func(string, domain.Property) (message.Message, error)
}

func New(store store.Store, notifier notifierproto.Client, timeNow func() time.Time) recyclingservicesproto.Handler {
	return &handler{
		store: store,
		fetcher: services.WebScraper(
			&http.Client{Timeout: time.Second * 30},
			services.ParseHTML,
			"https://recyclingservices.bromley.gov.uk/property",
		),
		notifier:        notifier,
		timeNow:         timeNow,
		propertyMessage: propertyMessageFunc(timeNow),
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

// SyncProperty will update the data stored for the property. This requires
// first fetching the latest property data (with the fetcher), then updating the
// store with the new content.
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

// NotifyProperty is used to send messages regarding the referenced property.
func (h *handler) NotifyProperty(ctx context.Context, body *recyclingservicesproto.NotifyPropertyRequest) (*recyclingservicesproto.NotifyPropertyResponse, error) {
	property, err := h.store.ReadProperty(ctx, body.PropertyId)
	if err != nil {
		return nil, err
	}
	msg, err := h.propertyMessage(body.MessageType, *property)
	if err != nil {
		return nil, err
	}
	if _, ok := msg.(message.NotSendable); ok {
		return &recyclingservicesproto.NotifyPropertyResponse{}, nil
	}
	if _, err := h.notifier.Notify(ctx, &notifierproto.NotifyRequest{
		Notifier: body.Notifier,
		Message:  msg.ToProto(),
	}); err != nil {
		return nil, err
	}
	return &recyclingservicesproto.NotifyPropertyResponse{}, nil
}

func propertyMessageFunc(timeNow func() time.Time) func(string, domain.Property) (message.Message, error) {
	servicesTomorrow := message.ServicesTomorrow(timeNow)
	servicesNextWeek := message.ServicesNextWeek(timeNow)
	describeProperty := message.DescribeProperty()
	return func(typ string, property domain.Property) (message.Message, error) {
		switch typ {
		case recyclingservicesproto.MessageTypeServicesTomorrow:
			return servicesTomorrow(property)
		case recyclingservicesproto.MessageTypeServicesNextWeek:
			return servicesNextWeek(property)
		case recyclingservicesproto.MessageTypeDescribeProperty:
			return describeProperty(property)
		default:
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("message type '%s' unsupported", typ))
		}
	}
}
