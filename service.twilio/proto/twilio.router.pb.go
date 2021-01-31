// DO NOT EDIT: Router was generated from 'twilio.proto'
package twilioproto

import (
	"context"
	"encoding/json"

	router "github.com/edstell/lambda-router"
	"google.golang.org/protobuf/encoding/protojson"
)

// If your request types implement the validator interface they will be 
// validated before being handed off to the handler for processing. 
type validator interface {
	Validate() error
}

// Handler is the exported interface you should implement to handle requests in
// your service.
type Handler interface {
	SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error)
}

// Router wraps the 'lambda-router', you should pass Router.Handle to 
// lambda.Start when initializing your lambda function.
type Router struct {
	*router.Router
}

// NewRouter initializes a router which will route requests to the handler
// provided. It will use the marshaler provided to pack errors for transport.
func NewRouter(handler Handler, marshaler func(error) (json.RawMessage, error)) *Router {
	router := router.New(router.MarshalErrorsWith(marshaler))
	router.Route("SendSMS", sendsms(handler.SendSMS))
	return &Router{
		Router: router,
	}
}

func sendsms(handler func(context.Context, *SendSMSRequest) (*SendSMSResponse, error)) router.Handler {
	return router.HandlerFunc(func(ctx context.Context, req json.RawMessage) (json.RawMessage, error) {
		body := &SendSMSRequest{}
		if err := protojson.Unmarshal(req, body); err != nil {
			return nil, err
		}
		var b interface{} = body
		if v, ok := b.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, err
			}
		}
		rsp, err := handler(ctx, body)
		if err != nil {
			return nil, err
		}
		bytes, err := protojson.Marshal(rsp)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	})
}
