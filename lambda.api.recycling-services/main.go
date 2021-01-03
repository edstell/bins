package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	svc "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/edstell/lambda/lambda.api.recycling-services/domain"
	"github.com/edstell/lambda/lambda.api.recycling-services/handler"
	recyclingservices "github.com/edstell/lambda/lambda.recycling-services/rpc"
	"github.com/edstell/lambda/libraries/api"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	logic := domain.NewBusinessLogic(recyclingservices.NewClient(svc.New(sess)))
	router := api.NewRouter()
	router.Route("GET", "/properties/{property}", handler.GETProperty(logic))
	lambda.Start(router.Handler)
}
