package domain

import (
	"context"

	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
	"github.com/edstell/lambda/lambda.recycling-services/store"
)

type Logic interface {
	ReadProperty(context.Context, string) (*recyclingservices.Property, error)
	WriteProperty(context.Context, recyclingservices.Property) error
}

type logic struct {
	store.Store
}

func NewLogic(store store.Store) Logic {
	return &logic{
		Store: store,
	}
}
