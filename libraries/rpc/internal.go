package rpc

import "encoding/json"

type request struct {
	ProcedureName string          `json:"procedure_name"`
	Body          json.RawMessage `json:"body"`
}

type response struct {
	Body json.RawMessage `json:"body"`
}
