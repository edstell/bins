package rpc

type request struct {
	ProcedureName string `json:"procedure_name"`
	Body          string `json:"body"`
}

type response struct {
	Body string `json:"body"`
}
