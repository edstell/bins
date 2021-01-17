package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestMissingString(t *testing.T) {
	t.Parallel()
	msg := &String{}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingMap(t *testing.T) {
	t.Parallel()
	msg := &Map{}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingList(t *testing.T) {
	t.Parallel()
	msg := &List{}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingBytes(t *testing.T) {
	t.Parallel()
	msg := &Bytes{}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingMessage(t *testing.T) {
	t.Parallel()
	msg := &Message{}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingOneof(t *testing.T) {
	t.Parallel()
	msg := &OneOf{}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingFieldInMapValue(t *testing.T) {
	t.Parallel()
	msg := &MapMessage{
		Map: map[string]*Message{
			"value": {},
		},
	}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingFieldInListValue(t *testing.T) {
	t.Parallel()
	msg := &ListMessage{
		List: []*Message{{}},
	}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingFieldInMessageValue(t *testing.T) {
	t.Parallel()
	msg := &MessageMessage{
		Msg: &Message{},
	}
	err := Validate(msg)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}

func TestMissingNestedRequiredField(t *testing.T) {
	t.Parallel()
	msg := &MessageMessage{}
	err := Validate(msg)
	require.NoError(t, err)
}
