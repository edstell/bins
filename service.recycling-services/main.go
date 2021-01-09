package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	svc "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/edstell/lambda/service.recycling-services/handler"
	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
	"github.com/edstell/lambda/service.recycling-services/store"
	twilio "github.com/edstell/lambda/service.twilio/rpc"
)

func timeNowUTC() time.Time {
	return time.Now().UTC()
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	store := store.NewDynamoDB(
		dynamodb.New(sess),
		timeNowUTC,
	)
	lambdaService := svc.New(sess)
	// Instrument the lambda client.
	xray.AWS(lambdaService.Client)
	client := twilio.NewClient(lambdaService, os.Getenv("TWILIO_ARN"))
	handler := handler.New(store, client, timeNowUTC)
	router := recyclingservices.NewRouter(handler)
	lambda.Start(router.Handler)
}
