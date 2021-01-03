package recyclingservices

import (
	"time"

	"github.com/edstell/lambda/libraries/rpc"
)

var Prodedures = []rpc.Procedure{
	{
		Name:     "ReadProperty",
		Request:  ReadPropertyRequest{},
		Response: ReadPropertyResponse{},
	},
	{
		Name:     "WriteProperty",
		Request:  WritePropertyRequest{},
		Response: WritePropertyResponse{},
	},
}

type Property struct {
	ID        string    `json:"property_id"`
	Services  []Service `json:"services"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Service struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Schedule    string    `json:"schedule"`
	LastService time.Time `json:"last_service"`
	NextService time.Time `json:"next_service"`
}

type ReadPropertyRequest struct {
	PropertyID string `json:"property_id"`
}

type ReadPropertyResponse struct {
	Property Property `json:"property"`
}

type WritePropertyRequest struct {
	PropertyID string    `json:"property_id"`
	Services   []Service `json:"services"`
}

type WritePropertyResponse struct {
	Property Property `json:"property"`
}
