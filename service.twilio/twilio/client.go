package twilio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
		return errors.MissingParam("To")
	}
	if _, ok := params["Body"]; !ok {
		return errors.MissingParam("Body")
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
		if message, ok := body["message"].(string); ok {
			return errors.NewKnown(rsp.StatusCode, message)
		}
		return errors.NewKnown(rsp.StatusCode, "failed to send sms with code")
	}

	return nil
}
