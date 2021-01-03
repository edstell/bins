package handler

import (
	"context"
	"time"

	"github.com/edstell/lambda/lambda.recycling-services/domain"
	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
)

type handler struct {
	bl      domain.BusinessLogic
	timeNow func() time.Time
}

func New(bl domain.BusinessLogic, timeNow func() time.Time) recyclingservices.Handler {
	return &handler{
		bl: bl,
	}
}

func (h *handler) ReadProperty(ctx context.Context, req recyclingservices.ReadPropertyRequest) (*recyclingservices.ReadPropertyResponse, error) {
	property, err := h.bl.ReadProperty(ctx, req.PropertyID)
	if err != nil {
		return nil, err
	}
	return &recyclingservices.ReadPropertyResponse{
		Property: *property,
	}, nil
}

func (h *handler) WriteProperty(ctx context.Context, req recyclingservices.WritePropertyRequest) (*recyclingservices.WritePropertyResponse, error) {
	property := recyclingservices.Property{
		ID:        req.PropertyID,
		Services:  req.Services,
		UpdatedAt: h.timeNow(),
	}
	if err := h.bl.WriteProperty(ctx, property); err != nil {
		return nil, err
	}
	return &recyclingservices.WritePropertyResponse{
		Property: property,
	}, nil
}
