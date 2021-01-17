package notifier

import (
	"context"

	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	twilioproto "github.com/edstell/lambda/service.twilio/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Notifier implementations notify by sending the passed message when called.
type Notifier interface {
	Notify(context.Context, Message) error
}

type NotifierFunc func(context.Context, Message) error

func (f NotifierFunc) Notify(ctx context.Context, message Message) error {
	return f(ctx, message)
}

// FromProtoFunc returns the function which should be used to marshal the proto
// Notifier implementation to a Notifier.
func FromProtoFunc(twilio *twilioproto.Client) func(*recyclingservicesproto.Notifier) Notifier {
	return func(notifier *recyclingservicesproto.Notifier) Notifier {
		switch v := notifier.Notifier.(type) {
		case *recyclingservicesproto.Notifier_Sms:
			return SMS(twilio, v.Sms.PhoneNumber)
		default:
			return NotifierFunc(func(context.Context, Message) error {
				return status.Errorf(codes.InvalidArgument, "unsupported notifier type '%T'", v)
			})
		}
	}
}

// SMS will send message via sms to the phone numer when called.
func SMS(twilio *twilioproto.Client, phoneNumber string) Notifier {
	return NotifierFunc(func(ctx context.Context, message Message) error {
		raw, err := message.Raw()
		if err != nil {
			return err
		}
		if _, err := twilio.SendSMS(ctx, &twilioproto.SendSMSRequest{
			PhoneNumber: phoneNumber,
			Message:     string(raw),
		}); err != nil {
			return err
		}
		return nil
	})
}
