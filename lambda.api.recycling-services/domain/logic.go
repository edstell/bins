package domain

import (
	"context"

	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
)

type BusinessLogic interface {
	ReadProperty(context.Context, string) (*recyclingservices.Property, error)
}

type logic struct {
	client *recyclingservices.Client
}

func NewBusinessLogic(client *recyclingservices.Client) BusinessLogic {
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
