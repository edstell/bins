package errors

import (
	"fmt"
	"net/http"
)

type Known interface {
	StatusCode() int
}

type known struct {
	statusCode int
	message    string
}

func (c *known) StatusCode() int {
	return c.statusCode
}

func (c *known) Error() string {
	return c.message
}

func NewKnown(statusCode int, message string) error {
	return &known{
		statusCode: statusCode,
		message:    message,
	}
}

func NotFound(resource string) error {
	return &known{
		statusCode: http.StatusNotFound,
		message:    fmt.Sprintf("'%s' was not found"),
	}
}

func BadRequest(reason string) error {
	return &known{
		statusCode: http.StatusBadRequest,
		message:    fmt.Sprintf("bad request: %s", reason),
	}
}

func MissingParam(param string) error {
	return BadRequest(fmt.Sprintf("missing param: %s", param))
}
