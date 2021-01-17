package notifier

import (
	"context"

	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	twilioproto "github.com/edstell/lambda/service.twilio/proto"
)

type Notifier interface {
	Notify(context.Context, Message) error
}

type NotifierFunc func(context.Context, Message) error

func (f NotifierFunc) Notify(ctx context.Context, message Message) error {
	return f(ctx, message)
}

func New(notifier *recyclingservicesproto.Notifier) Notifier {
	switch t := notifier.Notifier.(type) {
	case *recyclingservicesproto.Notifier_Sms:
	default:
		return NotifierFunc(func(context.Context, Message) error {
			return nil
		})
	}
}

func SMS(twilio *twilioproto.Client, phoneNumber string) Notifier {
	return NotifierFunc(func(ctx context.Context, message Message) error {
		body, err := message.Format()
		if err != nil {
			return err
		}
		if _, err := twilio.SendSMS(ctx, &twilioproto.SendSMSRequest{
			PhoneNumber: phoneNumber,
			Message:     body,
		}); err != nil {
			return err
		}
		return nil
	})
}
