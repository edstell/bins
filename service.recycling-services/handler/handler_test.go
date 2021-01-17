package handler

import (
	"testing"

	"github.com/edstell/lambda/libraries/validation"
	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"github.com/stretchr/testify/require"
)

func TestNotifyProperty(t *testing.T) {
	t.Parallel()
	body := &recyclingservicesproto.NotifyPropertyRequest{
		PropertyId:  "abc",
		MessageName: "name",
		Notifier: &recyclingservicesproto.NotifyPropertyRequest_Sms{
			Sms: &recyclingservicesproto.NotifyPropertyRequest_SMS{
				PhoneNumber: "+447...",
			},
		},
	}
	require.NoError(t, validation.Validate(body))
}
