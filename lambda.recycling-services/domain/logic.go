package domain

import (
	"context"

	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
	"github.com/edstell/lambda/lambda.recycling-services/store"
)

type BusinessLogic interface {
	ReadProperty(context.Context, string) (*recyclingservices.Property, error)
	WriteProperty(context.Context, recyclingservices.Property) error
}

type logic struct {
	store.Store
}

func NewBusinessLogic(store store.Store) BusinessLogic {
	return &logic{
		Store: store,
	}
}
