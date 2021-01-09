package recyclingservices

import (
	"context"
	"encoding/json"

	"github.com/edstell/lambda/libraries/rpc"
)

type Client struct {
	sendSMS rpc.Invoker
}

func NewClient(i rpc.LambdaInvoker, arn string) *Client {
	return &Client{
		sendSMS: rpc.Client(i, arn, "SendSMS"),
	}
}

func (c *Client) SendSMS(ctx context.Context, req SendSMSRequest) (*SendSMSResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	rsp, err := c.sendSMS.Invoke(ctx, payload)
	if err != nil {
		return nil, err
	}

	out := &SendSMSResponse{}
	if err := json.Unmarshal(rsp, out); err != nil {
		return nil, err
	}

	return out, nil
}
