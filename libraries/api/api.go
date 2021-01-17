package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return failed(status.Errorf(codes.PermissionDenied, "resource (%s) unavailable", req.Resource))
	}
	handler, ok := route[req.HTTPMethod]
	if !ok {
		return failed(status.Errorf(codes.PermissionDenied, "method (%s) unavailable for resource", req.HTTPMethod))
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
	httpStatus := http.StatusInternalServerError
	st, ok := status.FromError(err)
	if ok {
		httpStatus = HTTPStatusFromCode(st.Code())
	}
	return &Response{
		StatusCode: httpStatus,
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

// NOTE: Copied from https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}

func CodeFromHTTPStatus(httpStatus int) codes.Code {
	switch httpStatus {
	case http.StatusOK:
		return codes.OK
	case http.StatusRequestTimeout:
		return codes.Canceled
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	}
	return http.StatusInternalServerError
}
