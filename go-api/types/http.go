package types

import "encoding/json"

type WSRequest struct {
	ID     string          `validate:"required" json:"id"`
	Action string          `validate:"required" json:"action"`
	Token  string          `json:"token"`
	Data   json.RawMessage `validate:"required" json:"data"`
}

type WSResponse struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data,omitempty"`
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
}
