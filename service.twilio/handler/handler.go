package handler

import (
	"context"

	twilioproto "github.com/edstell/lambda/service.twilio/proto"
	"github.com/edstell/lambda/service.twilio/twilio"
)

type handler struct {
	client *twilio.Client
}

func New(client *twilio.Client) twilioproto.Handler {
	return &handler{
		client: client,
	}
}

func (h *handler) SendSMS(ctx context.Context, body *twilioproto.SendSMSRequest) (*twilioproto.SendSMSResponse, error) {
	if err := h.client.SendSMS(ctx, map[string]string{
		"To":   body.PhoneNumber,
		"Body": body.Message,
	}); err != nil {
		return nil, err
	}
	return &twilioproto.SendSMSResponse{}, nil
}
