package recyclingservices

import (
	"github.com/edstell/lambda/libraries/rpc"
)

var Prodedures = []rpc.Procedure{
	{
		Name:     "SendSMS",
		Request:  SendSMSRequest{},
		Response: SendSMSResponse{},
	},
}

type SendSMSRequest struct {
	To      string `json:"to"` // Destination phone number.
	Message string `json:"message"`
}

type SendSMSResponse struct {
}
