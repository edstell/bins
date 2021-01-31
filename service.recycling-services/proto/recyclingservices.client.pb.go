// DO NOT EDIT: Client was generated from 'recyclingservices.proto'
package recyclingservicesproto

import (
	"context"
	"encoding/json"

	awsreq "github.com/aws/aws-sdk-go/aws/request"
	invoker "github.com/edstell/lambda-invoker"
	"google.golang.org/protobuf/encoding/protojson"
)

// Client is a generated client for the service defined in 'RecyclingServices'.
// It exposes methods to call procedures available on the service, and handles
// the packing and unpacking of requests and responses for transport.
type Client interface {
	ReadProperty(context.Context, *ReadPropertyRequest, ...awsreq.Option) (*ReadPropertyResponse, error)
	SyncProperty(context.Context, *SyncPropertyRequest, ...awsreq.Option) (*SyncPropertyResponse, error)
	NotifyProperty(context.Context, *NotifyPropertyRequest, ...awsreq.Option) (*NotifyPropertyResponse, error)
}

// client is the internal implementation of Client.
type client struct {
	readproperty   *invoker.Invoker
	syncproperty   *invoker.Invoker
	notifyproperty   *invoker.Invoker
}

// NewClient initializes a Client, configuring it to use the provided 
// unmarshaler for unpacking errors to the error implementation of your choice.
func NewClient(i invoker.LambdaInvoker, arn string, unmarshaler func(json.RawMessage) error) Client {
	return &client{
		readproperty:   invoker.New(i, arn, invoker.AsProcedure("ReadProperty", unmarshaler)),
		syncproperty:   invoker.New(i, arn, invoker.AsProcedure("SyncProperty", unmarshaler)),
		notifyproperty:   invoker.New(i, arn, invoker.AsProcedure("NotifyProperty", unmarshaler)),
	}
}

func (c *client) ReadProperty(ctx context.Context, req *ReadPropertyRequest, opts ...awsreq.Option) (*ReadPropertyResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.readproperty.Invoke(ctx, payload, opts...)
	if err != nil {
		return nil, err
	}
	out := &ReadPropertyResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SyncProperty(ctx context.Context, req *SyncPropertyRequest, opts ...awsreq.Option) (*SyncPropertyResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.syncproperty.Invoke(ctx, payload, opts...)
	if err != nil {
		return nil, err
	}
	out := &SyncPropertyResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) NotifyProperty(ctx context.Context, req *NotifyPropertyRequest, opts ...awsreq.Option) (*NotifyPropertyResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.notifyproperty.Invoke(ctx, payload, opts...)
	if err != nil {
		return nil, err
	}
	out := &NotifyPropertyResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}
	return out, nil
}
