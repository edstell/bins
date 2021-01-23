package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	svc "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/edstell/lambda/service.notifier/handler"
	notifierproto "github.com/edstell/lambda/service.notifier/proto"
	twilioproto "github.com/edstell/lambda/service.twilio/proto"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	lambdaService := svc.New(sess)
	// Instrument the lambda client.
	xray.AWS(lambdaService.Client)
	twilio := twilioproto.NewClient(lambdaService, os.Getenv("TWILIO_ARN"))
	handler := handler.New(twilio)
	router := notifierproto.NewRouter(handler)
	lambda.Start(router.Handler)
}
