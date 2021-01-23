package handler

import (
	"context"

	"github.com/edstell/lambda/libraries/api"
	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
)

func GETProperty(recyclingServices recyclingservicesproto.Client) api.Handler {
	return func(ctx context.Context, req api.Request) (*api.Response, error) {
		rsp, err := recyclingServices.ReadProperty(ctx, &recyclingservicesproto.ReadPropertyRequest{
			PropertyId: req.PathParameters["property"],
		})
		if err != nil {
			return nil, err
		}
		return api.OK(rsp.Property)
	}
}
