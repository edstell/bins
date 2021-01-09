package recyclingservices

import "github.com/edstell/lambda/libraries/errors"

func (r SendSMSRequest) Validate() error {
	if r.To == "" {
		return errors.MissingParam("to")
	}
	if r.Message == "" {
		return errors.MissingParam("message")
	}
	return nil
}
