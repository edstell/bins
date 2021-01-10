package store

import (
	"context"

	"github.com/edstell/lambda/service.recycling-services/domain"
)

type Store interface {
	ReadProperty(context.Context, string) (*domain.Property, error)
	WriteProperty(context.Context, string, []domain.Service) (*domain.Property, error)
}
