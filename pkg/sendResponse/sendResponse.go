package sendResponse

import "github.com/gin-gonic/gin"

type response struct {
	Status string      `json:"status"`
	Msg    string      `json:"message"`
	Data   interface{} `json:"data"`
}

func Send(
	ctx *gin.Context,
	httpStatusCode int,
	status string, // "error" | "success"
	msg string,
	data interface{},
) {
	resp := response{
		Status: status,
		Msg:    msg,
		Data:   data,
	}

	ctx.JSON(httpStatusCode, resp)
}
