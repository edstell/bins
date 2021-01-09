package recyclingservices

import (
	"context"
	"encoding/json"

	"github.com/edstell/lambda/libraries/rpc"
)

type Handler interface {
	SendSMS(context.Context, SendSMSRequest) (*SendSMSResponse, error)
}

type Router struct {
	*rpc.Router
}

func NewRouter(handler Handler) *Router {
	router := rpc.NewRouter()
	router.Route("SendSMS", sendSMS(handler.SendSMS))
	return &Router{
		Router: router,
	}
}

func sendSMS(handler func(context.Context, SendSMSRequest) (*SendSMSResponse, error)) rpc.Handler {
	return func(ctx context.Context, req rpc.Request) (*rpc.Response, error) {
		body := &SendSMSRequest{}
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
