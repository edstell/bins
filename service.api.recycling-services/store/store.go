package store

import (
	"context"
	"time"

	"github.com/edstell/lambda/service.api.recycling-services/model"
)

type Property struct {
	ID        string          `json:"property_id"`
	Services  []model.Service `json:"services"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type Store interface {
	ReadProperty(context.Context, string) (*Property, error)
	WriteProperty(context.Context, Property) error
}
