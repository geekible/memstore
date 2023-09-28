package models

type Action int

const (
	_ Action = iota
	GET
	SET
	DELETE
	RESPONSE
	RESPONSE_ERROR
)

func (a *Action) String() string {
	switch *a {
	case GET:
		return "get"
	case SET:
		return "set"
	case DELETE:
		return "delete"
	case RESPONSE:
		return "response"
	case RESPONSE_ERROR:
		return "response_error"
	default:
		return "undefined"
	}
}

type TcpEnvelope struct {
	Action   Action `json:"action"`
	Key      string `json:"key,omitempty"`
	KeyValue struct {
		Key   string
		Value any
	} `json:"key_value,omitempty"`
	Response string `json:"response,omitempty"`
}
