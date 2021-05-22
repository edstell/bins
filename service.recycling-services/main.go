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
	"github.com/edstell/bins/libraries/errors"
	notifierproto "github.com/edstell/bins/service.notifier/proto"
	"github.com/edstell/bins/service.recycling-services/handler"
	recyclingservicesproto "github.com/edstell/bins/service.recycling-services/proto"
	"github.com/edstell/bins/service.recycling-services/store"
)

//go:generate go-bindata -o ./services/assets/assets.go -pkg assets -ignore=\.go$ -ignore=.DS_Store ./services/assets
//go:generate go fmt ./services/assets/assets.go

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
	notifier := notifierproto.NewClient(lambdaService, os.Getenv("NOTIFIER_ARN"), errors.Unmarshal)
	handler := handler.New(store, notifier, timeNowUTC)
	router := recyclingservicesproto.NewRouter(handler, errors.Marshal)
	lambda.Start(router.Handle)
}
