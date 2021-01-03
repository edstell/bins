package model

import (
	"time"

	"github.com/edstell/lambda/libraries/errors"
)

type Service struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Schedule    string    `json:"schedule"`
	LastService time.Time `json:"last_service"`
	NextService time.Time `json:"next_service"`
}

func (s Service) Validate() error {
	if s.Name == "" {
		return errors.MissingParam("name")
	}
	if s.Status == "" {
		return errors.MissingParam("status")
	}
	if s.Schedule == "" {
		return errors.MissingParam("schedule")
	}
	if s.LastService.IsZero() {
		return errors.MissingParam("last_service")
	}
	if s.NextService.IsZero() {
		return errors.MissingParam("next_service")
	}
	return nil
}
