package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edstell/lambda/service.twilio/handler"
	twilioproto "github.com/edstell/lambda/service.twilio/proto"
	"github.com/edstell/lambda/service.twilio/twilio"
)

func timeNowUTC() time.Time {
	return time.Now().UTC()
}

func main() {
	client := twilio.NewClient(
		twilio.WithSID(os.Getenv("SID")),
		twilio.WithAuthToken(os.Getenv("AUTH_TOKEN")),
		twilio.WithFrom(os.Getenv("FROM_NUMBER")),
	)
	handler := handler.New(client)
	router := twilioproto.NewRouter(handler)
	lambda.Start(router.Handler)
}
