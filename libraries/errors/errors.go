package errors

import (
	"encoding/json"
	"fmt"

	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// Marshal marshals the error as a grpc Status for transport. If the passed
// error isn't already a *grpc.Status, one is created from it initialized with
// the code 'internal' and err.Error() as the message.
func Marshal(err error) (json.RawMessage, error) {
	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Internal, err.Error())
	}
	bytes, err := protojson.Marshal(st.Proto())
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal always returns a grpc.Status error, attempting to unpack to a
// grpc.Status, falling back to json.Unmarshal and finally casting the raw bytes
// if all other decoding fails.
func Unmarshal(m json.RawMessage) error {
	st := &spb.Status{}
	if err := protojson.Unmarshal(m, st); err != nil {
		var i interface{}
		if err := json.Unmarshal(m, &i); err != nil {
			return status.Error(codes.Internal, fmt.Sprint(i))
		}
		return status.Error(codes.Internal, string(m))
	}
	return status.ErrorProto(st)
}
