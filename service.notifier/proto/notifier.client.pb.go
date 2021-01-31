// DO NOT EDIT: Client was generated from 'notifier.proto'
package notifierproto

import (
	"context"
	"encoding/json"

	awsreq "github.com/aws/aws-sdk-go/aws/request"
	invoker "github.com/edstell/lambda-invoker"
	"google.golang.org/protobuf/encoding/protojson"
)

// Client is a generated client for the service defined in 'NotifierSvc'.
// It exposes methods to call procedures available on the service, and handles
// the packing and unpacking of requests and responses for transport.
type Client interface {
	Notify(context.Context, *NotifyRequest, ...awsreq.Option) (*NotifyResponse, error)
}

// client is the internal implementation of Client.
type client struct {
	notify   *invoker.Invoker
}

// NewClient initializes a Client, configuring it to use the provided 
// unmarshaler for unpacking errors to the error implementation of your choice.
func NewClient(i invoker.LambdaInvoker, arn string, unmarshaler func(json.RawMessage) error) Client {
	return &client{
		notify:   invoker.New(i, arn, invoker.AsProcedure("Notify", unmarshaler)),
	}
}

func (c *client) Notify(ctx context.Context, req *NotifyRequest, opts ...awsreq.Option) (*NotifyResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.notify.Invoke(ctx, payload, opts...)
	if err != nil {
		return nil, err
	}
	out := &NotifyResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}
	return out, nil
}
