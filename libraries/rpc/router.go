package rpc

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Request struct {
	Body json.RawMessage
}

type Response struct {
	Body json.RawMessage
}

type Handler func(context.Context, Request) (*Response, error)

type Router struct {
	routes map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		routes: map[string]Handler{},
	}
}

func (r *Router) Route(procedureName string, handler Handler) {
	r.routes[procedureName] = handler
}

func (r *Router) Handler(ctx context.Context, req request) (*response, error) {
	handler, ok := r.routes[req.ProcedureName]
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "unsupported procedure '%s'", req.ProcedureName)
	}
	rsp, err := handler(ctx, Request{Body: []byte(req.Body)})
	if err != nil {
		return nil, err
	}
	return &response{Body: string(rsp.Body)}, err
}
