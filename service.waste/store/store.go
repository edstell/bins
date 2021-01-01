package store

import (
	"context"

	"github.com/edstell/waste-lambda/model"
)

type Store interface {
	ReadProperty(context.Context, string) ([]model.Service, error)
}
