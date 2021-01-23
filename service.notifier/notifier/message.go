package notifier

import (
	notifierproto "github.com/edstell/lambda/service.notifier/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Message implementations wrap the content of a message which can be sent with
// a Notifier.
type Message interface {
	Raw() []byte
}

type MessageFunc func() []byte

func (f MessageFunc) Raw() []byte {
	return f()
}

// MessageFromProto will unmarshal the message to a concrete internal
// implementation.
func MessageFromProto(message *notifierproto.Message) (Message, error) {
	switch v := message.Message.(type) {
	case *notifierproto.Message_BodyOnly_:
		return &BodyOnly{v.BodyOnly.Body}, nil
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unsupported message type '%T'", v)
	}
}
