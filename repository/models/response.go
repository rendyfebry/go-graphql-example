package models

// Response ...
type Response struct {
	Error   interface{} `json:"error,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
