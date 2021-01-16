// DO NOT EDIT: Router was autogenerated from 'github.com/edstell/lambda/service.recycling-services/proto/recyclingservices.proto'
package recyclingservicesproto

import (
	"context"

	"github.com/edstell/lambda/libraries/rpc"
	"github.com/edstell/lambda/libraries/validation"
	"google.golang.org/protobuf/encoding/protojson"
)

type Handler interface {
	ReadProperty(context.Context, *ReadPropertyRequest) (*ReadPropertyResponse, error)
	SyncProperty(context.Context, *SyncPropertyRequest) (*SyncPropertyResponse, error)
	NotifyProperty(context.Context, *NotifyPropertyRequest) (*NotifyPropertyResponse, error)
}

type Router struct {
	*rpc.Router
}

func NewRouter(handler Handler) *Router {
	router := rpc.NewRouter()
	router.Route("ReadProperty", readproperty(handler.ReadProperty))
	router.Route("SyncProperty", syncproperty(handler.SyncProperty))
	router.Route("NotifyProperty", notifyproperty(handler.NotifyProperty))
	return &Router{
		Router: router,
	}
}

func readproperty(handler func(context.Context, *ReadPropertyRequest) (*ReadPropertyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &ReadPropertyRequest{}
		if err := protojson.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		if err := validation.Validate(body); err != nil {
			return nil, err
		}

		rsp, err := handler(ctx, body)
		if err != nil {
			return nil, err
		}

		bytes, err := protojson.Marshal(rsp)
		if err != nil {
			return nil, err
		}

		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}

func syncproperty(handler func(context.Context, *SyncPropertyRequest) (*SyncPropertyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &SyncPropertyRequest{}
		if err := protojson.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		if err := validation.Validate(body); err != nil {
			return nil, err
		}

		rsp, err := handler(ctx, body)
		if err != nil {
			return nil, err
		}

		bytes, err := protojson.Marshal(rsp)
		if err != nil {
			return nil, err
		}

		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}

func notifyproperty(handler func(context.Context, *NotifyPropertyRequest) (*NotifyPropertyResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &NotifyPropertyRequest{}
		if err := protojson.Unmarshal(req.Body, body); err != nil {
			return nil, err
		}

		if err := validation.Validate(body); err != nil {
			return nil, err
		}

		rsp, err := handler(ctx, body)
		if err != nil {
			return nil, err
		}

		bytes, err := protojson.Marshal(rsp)
		if err != nil {
			return nil, err
		}

		return &rpc.Response{
			Body: bytes,
		}, nil
	}
}
