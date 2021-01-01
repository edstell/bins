package model

import "time"

type Service struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Schedule    string    `json:"schedule"`
	LastService time.Time `json:"last_service"`
	NextService time.Time `json:"next_service"`
}
