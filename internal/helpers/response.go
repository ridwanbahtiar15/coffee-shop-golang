package helpers

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GetResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}