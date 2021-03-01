package notifier

import (
	"context"

	notifierproto "github.com/edstell/bins/service.notifier/proto"
	twilioproto "github.com/edstell/bins/service.twilio/proto"
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
func FromProtoFunc(twilio twilioproto.Client) func(*notifierproto.Notifier) (Notifier, error) {
	return func(notifier *notifierproto.Notifier) (Notifier, error) {
		switch v := notifier.Notifier.(type) {
		case *notifierproto.Notifier_Sms:
			return SMS(twilio, v.Sms.PhoneNumber), nil
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unsupported notifier type '%T'", v)
		}
	}
}
