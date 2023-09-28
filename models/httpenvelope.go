package models

type HttpEnvelope struct {
	KeyValue struct {
		Key   string
		Value any
	} `json:"key_value,omitempty"`
	Response string `json:"response,omitempty"`
}
