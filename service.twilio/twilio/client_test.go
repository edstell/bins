package twilio

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendSMS(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	c := &Client{
		client:  server.Client(),
		baseURL: server.URL,
		sid:     "SID",
	}
	err := c.SendSMS(context.Background(), map[string]string{
		"to":   "+4412345678910",
		"from": "+4410987654321",
		"body": "test message",
	})
	require.NoError(t, err)
}
