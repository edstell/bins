package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/edstell/lambda/service.twilio/handler"
	rpc "github.com/edstell/lambda/service.twilio/rpc"
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
	router := rpc.NewRouter(handler)
	lambda.Start(router.Handler)
}
