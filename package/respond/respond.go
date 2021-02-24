package respond

type ResponseError struct {
	Messages interface{} `json:"messages"`
}

// ResponseSuccess ...
type ResponseSuccess struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func ErrorResponse(message interface{}) ResponseError {
	err := []interface{}{message}
	res := ResponseError{Messages: err}

	return res
}

// SuccessResponse ...
func SuccessResponse(data interface{}, meta interface{}) ResponseSuccess {
	return ResponseSuccess{
		Data: data,
		Meta: meta,
	}
}