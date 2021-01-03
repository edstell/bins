package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/edstell/lambda/lambda.api.recycling-services/domain"
	"github.com/edstell/lambda/lambda.api.recycling-services/model"
	"github.com/edstell/lambda/lambda.api.recycling-services/store"
	"github.com/edstell/lambda/libraries/api"
	"github.com/edstell/lambda/libraries/errors"
)

func GETProperty(logic domain.BusinessLogic) api.Handler {
	return func(ctx context.Context, req api.Request) (*api.Response, error) {
		property, err := logic.ReadProperty(ctx, req.PathParameters["property"])
		if err != nil {
			return nil, err
		}
		return api.OK(property)
	}
}

func PUTProperty(logic domain.BusinessLogic, timeNow func() time.Time) api.Handler {
	type request struct {
		Services []model.Service `json:"services"`
	}
	return func(ctx context.Context, req api.Request) (*api.Response, error) {
		body := &request{}
		if err := json.Unmarshal([]byte(req.Body), body); err != nil {
			return nil, errors.BadRequest(fmt.Sprintf("body malformed: %v", err))
		}
		if len(body.Services) == 0 {
			return nil, errors.MissingParam("services")
		}
		for _, service := range body.Services {
			if err := service.Validate(); err != nil {
				return nil, errors.BadRequest(fmt.Sprintf("service[s] malformed: %v", err))
			}
		}

		if err := logic.WriteProperty(ctx, store.Property{
			ID:        req.PathParameters["property"],
			Services:  body.Services,
			UpdatedAt: timeNow(),
		}); err != nil {
			return nil, err
		}

		return api.OK(nil)
	}
}
