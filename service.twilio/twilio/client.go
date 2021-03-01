package twilio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/edstell/bins/libraries/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	sid       string
	authToken string
	baseURL   string
	from      string
	client    *http.Client
}

func NewClient(options ...func(*Client)) *Client {
	client := &Client{
		baseURL: "https://api.twilio.com/2010-04-01",
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func WithSID(sid string) func(*Client) {
	return func(client *Client) {
		client.sid = sid
	}
}

func WithAuthToken(authToken string) func(*Client) {
	return func(client *Client) {
		client.authToken = authToken
	}
}

func WithFrom(from string) func(*Client) {
	return func(client *Client) {
		client.from = from
	}
}

func WithBaseURL(url string) func(*Client) {
	return func(client *Client) {
		client.baseURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(client *Client) {
		client.client = httpClient
	}
}

func (c *Client) SendSMS(ctx context.Context, params map[string]string) error {
	if _, ok := params["To"]; !ok {
		return status.Error(codes.InvalidArgument, "'To' must be provided")
	}
	if _, ok := params["Body"]; !ok {
		return status.Error(codes.InvalidArgument, "'Body' must be provided")
	}

	query := url.Values{}
	query.Add("From", c.from)
	for key, value := range params {
		query.Add(key, value)
	}
	reader := strings.NewReader(query.Encode())
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/Accounts/%s/Messages.json", c.baseURL, c.sid), reader)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.sid, c.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)

	rsp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode < 200 || rsp.StatusCode >= 300 {
		body := map[string]interface{}{}
		if err := json.NewDecoder(rsp.Body).Decode(&body); err != nil {
			return err
		}
		message := "failed to send sms with code"
		if m, ok := body["message"].(string); ok {
			message = m
		}
		return status.Error(api.CodeFromHTTPStatus(rsp.StatusCode), message)
	}

	return nil
}
