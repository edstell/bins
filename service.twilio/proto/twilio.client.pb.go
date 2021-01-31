// DO NOT EDIT: Client was generated from 'twilio.proto'
package twilioproto

import (
	"context"
	"encoding/json"

	awsreq "github.com/aws/aws-sdk-go/aws/request"
	invoker "github.com/edstell/lambda-invoker"
	"google.golang.org/protobuf/encoding/protojson"
)

// Client is a generated client for the service defined in 'Twilio'.
// It exposes methods to call procedures available on the service, and handles
// the packing and unpacking of requests and responses for transport.
type Client interface {
	SendSMS(context.Context, *SendSMSRequest, ...awsreq.Option) (*SendSMSResponse, error)
}

// client is the internal implementation of Client.
type client struct {
	sendsms   *invoker.Invoker
}

// NewClient initializes a Client, configuring it to use the provided 
// unmarshaler for unpacking errors to the error implementation of your choice.
func NewClient(i invoker.LambdaInvoker, arn string, unmarshaler func(json.RawMessage) error) Client {
	return &client{
		sendsms:   invoker.New(i, arn, invoker.AsProcedure("SendSMS", unmarshaler)),
	}
}

func (c *client) SendSMS(ctx context.Context, req *SendSMSRequest, opts ...awsreq.Option) (*SendSMSResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.sendsms.Invoke(ctx, payload, opts...)
	if err != nil {
		return nil, err
	}
	out := &SendSMSResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}
	return out, nil
}
