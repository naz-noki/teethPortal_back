package sendResponse

import "github.com/gin-gonic/gin"

type ResponseMeta struct {
	Page    int `json:"page"`
	Limit   int `json:"limit"`
	Counter int `json:"counter"`
}

type Response[T interface{}] struct {
	Status string        `json:"status"`
	Msg    string        `json:"message"`
	Meta   *ResponseMeta `json:"meta"`
	Data   T             `json:"data"`
}

func Send(
	ctx *gin.Context,
	httpStatusCode int,
	status string, // "error" | "success"
	msg string,
	page int,
	limit int,
	counter int,
	data interface{},
) {
	resp := Response[interface{}]{
		Status: status,
		Msg:    msg,
		Meta: &ResponseMeta{
			Page:    page,
			Limit:   limit,
			Counter: counter,
		},
		Data: data,
	}

	ctx.JSON(httpStatusCode, resp)
}
