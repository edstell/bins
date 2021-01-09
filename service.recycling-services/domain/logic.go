package domain

import (
	"context"
	"errors"
	"net/http"
	"time"

	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
	"github.com/edstell/lambda/service.recycling-services/services"
	"github.com/edstell/lambda/service.recycling-services/store"
)

type Logic interface {
	ReadProperty(context.Context, string) (*recyclingservices.Property, error)
	WriteProperty(context.Context, string, []recyclingservices.Service) (*recyclingservices.Property, error)
	SyncProperty(context.Context, string) (*recyclingservices.Property, error)
}

type logic struct {
	store.Store
	fetcher services.Fetcher
}

func NewLogic(store store.Store) Logic {
	return &logic{
		Store: store,
		fetcher: services.WebScraper(&http.Client{
			Timeout: time.Second * 30,
		},
			services.ParseHTML,
			"https://recyclingservices.bromley.gov.uk/property",
		),
	}
}

func (l *logic) SyncProperty(ctx context.Context, propertyID string) (*recyclingservices.Property, error) {
	services, err := l.fetcher.Fetch(ctx, propertyID)
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, errors.New("no services fetched")
	}
	return l.Store.WriteProperty(ctx, propertyID, services)
}
