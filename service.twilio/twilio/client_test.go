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
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	c := &Client{
		client:  server.Client(),
		baseURL: server.URL,
		sid:     "SID",
	}
	err := c.SendSMS(context.Background(), map[string]string{
		"To":   "+4412345678910",
		"Body": "test message",
	})
	require.NoError(t, err)
}
