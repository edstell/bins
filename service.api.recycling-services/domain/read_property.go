package domain

import (
	"context"

	"github.com/edstell/lambda/api"
	"github.com/edstell/lambda/service.api.recycling-services/store"
)

// ReadProperty retrieves the property referenced from persistent storage.
func ReadProperty(store store.Store) api.Handler {
	return func(ctx context.Context, req api.Request) (*api.Response, error) {
		services, err := store.ReadProperty(ctx, req.PathParameters["property"])
		if err != nil {
			return nil, err
		}
		return api.OK(services)
	}
}
