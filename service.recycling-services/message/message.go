package message

import (
	notifierproto "github.com/edstell/bins/service.notifier/proto"
)

// Message implementations describe a message which can be sent with
// service.notifier.
type Message interface {
	ToProto() *notifierproto.Message
}

type MessageFunc func() *notifierproto.Message

func (f MessageFunc) ToProto() *notifierproto.Message {
	return f()
}

// NotSendable indicates that a message is unsendable, messages which implement
// this shouldn't be sent (and will panic if 'ToProto' is called on them).
type NotSendable interface {
	NotSendable()
}

type dontSend struct {
	Message
}

var _ NotSendable = &dontSend{}

func (d *dontSend) NotSendable() {}

// DontSend returns a 'NotSendable' Message.
func DontSend() Message {
	return &dontSend{}
}

// BodyOnly is a Message implementation which marshals to the BodyOnly Message
// type.
type BodyOnly struct {
	Body string
}

func (b *BodyOnly) ToProto() *notifierproto.Message {
	return &notifierproto.Message{
		Message: &notifierproto.Message_BodyOnly_{
			BodyOnly: &notifierproto.Message_BodyOnly{
				Body: b.Body,
			},
		},
	}
}
