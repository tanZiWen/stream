package domain

type Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SuccessResponse(data interface{}) (r *Response) {
	return &Response{
		Code: "ok",
		Data: data,
	}
}

func SuccessPairResponse(key string, value interface{}) (r *Response) {
	return &Response{
		Code: "ok",
		Data: map[string]interface{}{key: value},
	}
}

func FailureResponse(code string, message string) (r *Response) {
	return &Response{
		Code: code,
		Message: message,
	}
}
