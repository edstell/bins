package store

import (
	"context"

	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
)

type Store interface {
	ReadProperty(context.Context, string) (*recyclingservicesproto.Property, error)
	WriteProperty(context.Context, string, []*recyclingservicesproto.Service) (*recyclingservicesproto.Property, error)
}
