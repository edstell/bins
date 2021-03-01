package handler

import (
	"context"

	"github.com/edstell/bins/libraries/api"
	recyclingservicesproto "github.com/edstell/bins/service.recycling-services/proto"
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
