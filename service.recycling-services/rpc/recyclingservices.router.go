package recyclingservices

import (
	"context"
	"encoding/json"

	"github.com/edstell/lambda/libraries/rpc"
)

type Handler interface {
	ReadProperty(context.Context, ReadPropertyRequest) (*ReadPropertyResponse, error)
	WriteProperty(context.Context, WritePropertyRequest) (*WritePropertyResponse, error)
	SyncProperty(context.Context, SyncPropertyRequest) (*SyncPropertyResponse, error)
	NotifyProperty(context.Context, NotifyPropertyRequest) (*NotifyPropertyResponse, error)
}

type Router struct {
	*rpc.Router
}

func NewRouter(handler Handler) *Router {
	router := rpc.NewRouter()
	router.Route("ReadProperty", readProperty(handler.ReadProperty))
	router.Route("WriteProperty", writeProperty(handler.WriteProperty))
	router.Route("SyncProperty", syncProperty(handler.SyncProperty))
	router.Route("NotifyProperty", notifyProperty(handler.NotifyProperty))
	return &Router{
		Router: router,
	}
}

func readProperty(handler func(context.Context, ReadPropertyRequest) (*ReadPropertyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &ReadPropertyRequest{}
		if err := json.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		rsp, err := handler(ctx, *body)
		if err != nil {
			return nil, err
		}

		bytes, err := json.Marshal(rsp)
		if err != nil {
			return nil, err
		}

		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}

func writeProperty(handler func(context.Context, WritePropertyRequest) (*WritePropertyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &WritePropertyRequest{}
		if err := json.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		rsp, err := handler(ctx, *body)
		if err != nil {
			return nil, err
		}

		bytes, err := json.Marshal(rsp)
		if err != nil {
			return nil, err
		}

		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}

func syncProperty(handler func(context.Context, SyncPropertyRequest) (*SyncPropertyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &SyncPropertyRequest{}
		if err := json.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		rsp, err := handler(ctx, *body)
		if err != nil {
			return nil, err
		}

		bytes, err := json.Marshal(rsp)
		if err != nil {
			return nil, err
		}

		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}

func notifyProperty(handler func(context.Context, NotifyPropertyRequest) (*NotifyPropertyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &NotifyPropertyRequest{}
		if err := json.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		rsp, err := handler(ctx, *body)
		if err != nil {
			return nil, err
		}

		bytes, err := json.Marshal(rsp)
		if err != nil {
			return nil, err
		}

		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}
