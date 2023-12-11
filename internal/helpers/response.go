package helpers

type Response struct {
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
	Meta    map[string]any `json:"meta,omitempty"`
}

func GetResponse(message string, data interface{}, meta map[string]any) Response {
	return Response{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}