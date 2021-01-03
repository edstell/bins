package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/edstell/lambda/libraries/errors"
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
		r.routes[resource] = route
	}
	route[method] = handler
}

func (r *Router) Handler(ctx context.Context, req Request) (*Response, error) {
	route, ok := r.routes[req.Resource]
	if !ok {
		return failed(errors.NewKnown(http.StatusBadRequest, fmt.Sprintf("invalid resource: %s", req.Resource)))
	}
	handler, ok := route[req.HTTPMethod]
	if !ok {
		return failed(errors.NewKnown(http.StatusMethodNotAllowed, "unsupported method for resource"))
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
	if k, ok := err.(errors.Known); ok {
		statusCode = k.StatusCode()
	}
	return &Response{
		StatusCode: statusCode,
		Body:       string(message),
	}, nil
}

func OK(body interface{}) (*Response, error) {
	rsp := ""
	if body != nil {
		bytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rsp = string(bytes)
	}
	return &Response{
		StatusCode: http.StatusOK,
		Body:       rsp,
	}, nil
}
