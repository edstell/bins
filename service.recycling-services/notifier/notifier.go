package notifier

import (
	"context"

	twilio "github.com/edstell/lambda/service.twilio/rpc"
)

type Notifier interface {
	Notify(context.Context, Message) error
}

type NotifierFunc func(context.Context, Message) error

func (f NotifierFunc) Notify(ctx context.Context, message Message) error {
	return f(ctx, message)
}

func SMS(client *twilio.Client, phoneNumber string) Notifier {
	return NotifierFunc(func(ctx context.Context, message Message) error {
		body, err := message.Format()
		if err != nil {
			return err
		}
		if _, err := client.SendSMS(ctx, twilio.SendSMSRequest{
			To:      phoneNumber,
			Message: body,
		}); err != nil {
			return err
		}
		return nil
	})
}
