package handler

import (
	"context"

	"github.com/edstell/lambda/libraries/api"
	"github.com/edstell/lambda/service.api.recycling-services/domain"
)

func GETProperty(logic domain.Logic) api.Handler {
	return func(ctx context.Context, req api.Request) (*api.Response, error) {
		property, err := logic.ReadProperty(ctx, req.PathParameters["property"])
		if err != nil {
			return nil, err
		}
		return api.OK(property)
	}
}
