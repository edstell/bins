package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/edstell/waste-lambda/errors"
)

type Handler func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

type route map[string]Handler

type Router struct {
	routes map[string]route
}

func New() *Router {
	return &Router{
		routes: map[string]route{},
	}
}

func (r *Router) Route(method, resource string, handler Handler) {
	route, ok := r.routes[resource]
	if !ok {
		route = map[string]Handler{}
	}
	route[method] = handler
}

func (r *Router) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	route, ok := r.routes[req.Resource]
	if !ok {
		return failed(errors.NewClient(http.StatusBadRequest, "invalid resource"))
	}
	handler, ok := route[req.HTTPMethod]
	if !ok {
		return failed(errors.NewClient(http.StatusMethodNotAllowed, "unsupported method for resource"))
	}
	rsp, err := handler(ctx, req)
	if err != nil {
		return failed(err)
	}
	return rsp, nil
}

func failed(err error) (*events.APIGatewayProxyResponse, error) {
	message, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
	if err != nil {
		return nil, err
	}
	statusCode := http.StatusInternalServerError
	if c, ok := err.(errors.Client); ok {
		statusCode = c.StatusCode()
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(message),
	}, nil
}

func OK(body interface{}) (*events.APIGatewayProxyResponse, error) {
	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bytes),
	}, nil
}
