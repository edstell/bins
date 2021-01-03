package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/edstell/lambda/lambda.api.recycling-services/domain"
	"github.com/edstell/lambda/lambda.api.recycling-services/handler"
	"github.com/edstell/lambda/lambda.api.recycling-services/store"
	"github.com/edstell/lambda/libraries/api"
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
	bl := domain.NewBusinessLogic(store)
	router := api.NewRouter()
	router.Route("GET", "/properties/{property}", handler.GETProperty(bl))
	router.Route("PUT", "/properties/{property}", handler.PUTProperty(bl, timeNowUTC))
	lambda.Start(router.Handler)
}
