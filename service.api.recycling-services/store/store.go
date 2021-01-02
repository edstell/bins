package store

import (
	"context"

	"github.com/edstell/lambda/service.api.recycling-services/model"
)

type Store interface {
	ReadProperty(context.Context, string) ([]model.Service, error)
}
