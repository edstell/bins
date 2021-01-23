package notifier

import (
	"context"

	twilioproto "github.com/edstell/lambda/service.twilio/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SMS will send the message's body as an sms.
func SMS(twilio twilioproto.Client, phoneNumber string) Notifier {
	return NotifierFunc(func(ctx context.Context, message Message) error {
		m, ok := message.(interface {
			Body() string
		})
		if !ok {
			return status.Error(codes.InvalidArgument, "sms notifier requires a message with a body")
		}
		if _, err := twilio.SendSMS(ctx, &twilioproto.SendSMSRequest{
			PhoneNumber: phoneNumber,
			Message:     m.Body(),
		}); err != nil {
			return err
		}
		return nil
	})
}
