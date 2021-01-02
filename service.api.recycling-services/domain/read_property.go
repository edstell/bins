package domain

import (
	"context"

	"github.com/edstell/lambda/api"
	"github.com/edstell/lambda/service.api.recycling-services/store"
)

// ReadProperty retrieves the property referenced from persistent storage.
func ReadProperty(store store.Store) api.Handler {
	type response struct {
		Method         string            `json:"method"`
		Resource       string            `json:"resource"`
		PathParameters map[string]string `json:"path_parameters"`
	}
	return func(ctx context.Context, req api.Request) (*api.Response, error) {
		return api.OK(response{
			Method:         req.HTTPMethod,
			Resource:       req.Resource,
			PathParameters: req.PathParameters,
		})
	}
}
