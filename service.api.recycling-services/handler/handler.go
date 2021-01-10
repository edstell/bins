package handler

import (
	"context"

	"github.com/edstell/lambda/libraries/api"
	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
)

func GETProperty(client *recyclingservices.Client) api.Handler {
	return func(ctx context.Context, req api.Request) (*api.Response, error) {
		rsp, err := client.ReadProperty(ctx, recyclingservices.ReadPropertyRequest{
			PropertyID: req.PathParameters["property"],
		})
		if err != nil {
			return nil, err
		}
		return api.OK(rsp.Property)
	}
}
