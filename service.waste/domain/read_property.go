package domain

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/edstell/waste-lambda/router"
	"github.com/edstell/waste-lambda/store"
)

// ReadProperty retrieves the property referenced from persistent storage.
func ReadProperty(store store.Store) router.Handler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		services, err := store.ReadProperty(ctx, req.PathParameters["property"])
		if err != nil {
			return nil, err
		}
		return router.OK(services)
	}
}
