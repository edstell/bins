package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/edstell/lambda/errors"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse

type Handler func(context.Context, Request) (*Response, error)

type route map[string]Handler

type Router struct {
	routes map[string]route
}

func NewRouter() *Router {
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

func (r *Router) Handler(ctx context.Context, req Request) (*Response, error) {
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

func failed(err error) (*Response, error) {
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
	return &Response{
		StatusCode: statusCode,
		Body:       string(message),
	}, nil
}

func OK(body interface{}) (*Response, error) {
	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: http.StatusOK,
		Body:       string(bytes),
	}, nil
}
