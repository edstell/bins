package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/edstell/lambda/service.recycling-services/domain"
	"github.com/edstell/lambda/service.recycling-services/handler"
	recyclingservices "github.com/edstell/lambda/service.recycling-services/rpc"
	"github.com/edstell/lambda/service.recycling-services/store"
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
	handler := handler.New(domain.NewLogic(store))
	router := recyclingservices.NewRouter(handler)
	lambda.Start(router.Handler)
}
