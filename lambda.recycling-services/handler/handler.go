package handler

import (
	"context"

	"github.com/edstell/lambda/lambda.recycling-services/domain"
	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
)

type handler struct {
	l domain.Logic
}

func New(l domain.Logic) recyclingservices.Handler {
	return &handler{
		l: l,
	}
}

func (h *handler) ReadProperty(ctx context.Context, body recyclingservices.ReadPropertyRequest) (*recyclingservices.ReadPropertyResponse, error) {
	property, err := h.l.ReadProperty(ctx, body.PropertyID)
	if err != nil {
		return nil, err
	}
	return &recyclingservices.ReadPropertyResponse{
		Property: *property,
	}, nil
}

func (h *handler) WriteProperty(ctx context.Context, body recyclingservices.WritePropertyRequest) (*recyclingservices.WritePropertyResponse, error) {
	property, err := h.l.WriteProperty(ctx, body.PropertyID, body.Services)
	if err != nil {
		return nil, err
	}
	return &recyclingservices.WritePropertyResponse{
		Property: *property,
	}, nil
}

func (h *handler) SyncProperty(ctx context.Context, body recyclingservices.SyncPropertyRequest) (*recyclingservices.SyncPropertyResponse, error) {
	property, err := h.l.SyncProperty(ctx, body.PropertyID)
	if err != nil {
		return nil, err
	}
	return &recyclingservices.SyncPropertyResponse{
		Property: *property,
	}, nil
}
