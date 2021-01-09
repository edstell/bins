package domain

import (
	"context"

	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
)

type Logic interface {
	ReadProperty(context.Context, string) (*recyclingservices.Property, error)
}

type logic struct {
	client *recyclingservices.Client
}

func NewLogic(client *recyclingservices.Client) Logic {
	return &logic{
		client: client,
	}
}

func (l *logic) ReadProperty(ctx context.Context, propertyID string) (*recyclingservices.Property, error) {
	rsp, err := l.client.ReadProperty(ctx, recyclingservices.ReadPropertyRequest{
		PropertyID: propertyID,
	})
	if err != nil {
		return nil, err
	}
	return &rsp.Property, nil
}
