package domain

import (
	"context"

	"github.com/edstell/lambda/service.api.recycling-services/store"
)

type BusinessLogic interface {
	ReadProperty(context.Context, string) (*store.Property, error)
	WriteProperty(context.Context, store.Property) error
}

type logic struct {
	store.Store
}

func NewBusinessLogic(store store.Store) BusinessLogic {
	return &logic{
		Store: store,
	}
}
