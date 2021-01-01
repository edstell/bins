package errors

type Client interface {
	StatusCode() int
}

type client struct {
	statusCode int
	message    string
}

func (c *client) StatusCode() int {
	return c.statusCode
}

func (c *client) Error() string {
	return c.message
}

func NewClient(statusCode int, message string) error {
	return &client{
		statusCode: statusCode,
		message:    message,
	}
}
