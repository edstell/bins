package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edstell/lambda/api"
	"github.com/edstell/lambda/service.api.recycling-services/domain"
)

func timeNowUTC() time.Time {
	return time.Now().UTC()
}

func main() {
	router := api.NewRouter()
	router.Route("GET", "/properties/{property}", domain.ReadProperty(nil))
	lambda.Start(router.Handler)
}
