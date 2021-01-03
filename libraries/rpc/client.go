package rpc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsreq "github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Procedure struct {
	Name     string
	Request  interface{}
	Response interface{}
}

type LambdaInvoker interface {
	InvokeWithContext(context.Context, *lambda.InvokeInput, ...awsreq.Option) (*lambda.InvokeOutput, error)
}

type Invoker interface {
	Invoke(context.Context, []byte) ([]byte, error)
}

type InvokerFunc func(context.Context, []byte) ([]byte, error)

func (f InvokerFunc) Invoke(ctx context.Context, bytes []byte) ([]byte, error) {
	return f(ctx, bytes)
}

func Client(invoker LambdaInvoker, arn, procedureName string) Invoker {
	return InvokerFunc(func(ctx context.Context, body []byte) ([]byte, error) {
		payload, err := json.Marshal(request{
			ProcedureName: procedureName,
			Body:          string(body),
		})
		if err != nil {
			return nil, err
		}

		output, err := invoker.InvokeWithContext(ctx, &lambda.InvokeInput{
			FunctionName:   aws.String(arn),
			InvocationType: aws.String("RequestResponse"),
			Payload:        payload,
		})
		if err != nil {
			return nil, err
		}

		if message := output.FunctionError; message != nil {
			return nil, errors.New(*message)
		}

		rsp := &response{}
		if err := json.Unmarshal(output.Payload, rsp); err != nil {
			return nil, err
		}

		return []byte(rsp.Body), nil
	})
}
