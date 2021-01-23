package handler

import (
	"context"
	"time"

	"github.com/edstell/lambda/service.notifier/notifier"
	notifierproto "github.com/edstell/lambda/service.notifier/proto"
	twilioproto "github.com/edstell/lambda/service.twilio/proto"
)

func timeNowUTC() time.Time {
	return time.Now().UTC()
}

type handler struct {
	notifierFromProto func(*notifierproto.Notifier) (notifier.Notifier, error)
	messageFromProto  func(*notifierproto.Message) (notifier.Message, error)
}

func New(twilio twilioproto.Client) notifierproto.Handler {
	return &handler{
		notifierFromProto: notifier.FromProtoFunc(twilio),
		messageFromProto:  notifier.MessageFromProto,
	}
}

// Notify will send the message using the notifier.
func (h *handler) Notify(ctx context.Context, body *notifierproto.NotifyRequest) (*notifierproto.NotifyResponse, error) {
	notifier, err := h.notifierFromProto(body.Notifier)
	if err != nil {
		return nil, err
	}
	message, err := h.messageFromProto(body.Message)
	if err != nil {
		return nil, err
	}
	if err := notifier.Notify(ctx, message); err != nil {
		return nil, err
	}
	return &notifierproto.NotifyResponse{}, nil
}
