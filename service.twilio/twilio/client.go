package twilio

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/edstell/lambda/libraries/errors"
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

func WithFrom(authToken string) func(*Client) {
	return func(client *Client) {
		client.authToken = authToken
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
	if _, ok := params["to"]; !ok {
		return errors.MissingParam("to")
	}
	if _, ok := params["body"]; !ok {
		return errors.MissingParam("body")
	}

	query := url.Values{}
	query.Add("from", c.from)
	for key, value := range params {
		query.Add(key, value)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/Accounts/%s/Messages.json?%s", c.baseURL, c.sid, query.Encode()), nil)
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
		return errors.NewKnown(rsp.StatusCode, "failed to post new message")
	}

	return nil
}
