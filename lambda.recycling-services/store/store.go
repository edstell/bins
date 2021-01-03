package store

import (
	"context"

	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
)

type Store interface {
	ReadProperty(context.Context, string) (*recyclingservices.Property, error)
	WriteProperty(context.Context, recyclingservices.Property) error
}
