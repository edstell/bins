package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edstell/waste-lambda/domain"
	"github.com/edstell/waste-lambda/router"
)

func timeNowUTC() time.Time {
	return time.Now().UTC()
}

func main() {
	router := router.New()
	router.Route("GET", "/properties/{property}", domain.ReadProperty(nil))
	lambda.Start(router.Handler)
}
