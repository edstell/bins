package recyclingservices

import (
	"context"
	"encoding/json"

	"github.com/edstell/lambda/libraries/rpc"
)

type Client struct {
	readProperty  rpc.Invoker
	writeProperty rpc.Invoker
}

func NewClient(i rpc.LambdaInvoker, arn string) *Client {
	return &Client{
		readProperty:  rpc.Client(i, arn, "ReadProperty"),
		writeProperty: rpc.Client(i, arn, "WriteProperty"),
	}
}

func (c *Client) ReadProperty(ctx context.Context, req ReadPropertyRequest) (*ReadPropertyResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.readProperty.Invoke(ctx, payload)
	if err != nil {
		return nil, err
	}

	out := &ReadPropertyResponse{}
	if err := json.Unmarshal(rsp, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) WriteProperty(ctx context.Context, req WritePropertyRequest) (*WritePropertyResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.writeProperty.Invoke(ctx, payload)
	if err != nil {
		return nil, err
	}

	out := &WritePropertyResponse{}
	if err := json.Unmarshal(rsp, out); err != nil {
		return nil, err
	}

	return out, nil
}
