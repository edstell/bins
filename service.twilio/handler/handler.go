package handler

import (
	"context"

	rpc "github.com/edstell/lambda/service.twilio/rpc"
	"github.com/edstell/lambda/service.twilio/twilio"
)

type handler struct {
	client *twilio.Client
}

func New(client *twilio.Client) rpc.Handler {
	return &handler{
		client: client,
	}
}

func (h *handler) SendSMS(ctx context.Context, body rpc.SendSMSRequest) (*rpc.SendSMSResponse, error) {
	if err := h.client.SendSMS(ctx, map[string]string{
		"To":   body.To,
		"Body": body.Message,
	}); err != nil {
		return nil, err
	}
	return &rpc.SendSMSResponse{}, nil
}
