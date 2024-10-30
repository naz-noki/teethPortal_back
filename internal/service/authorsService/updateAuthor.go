package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (as *authorsService) UpdateAuthor(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.SaveAuthorBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("UpdateAuthor: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}
	// Получаем id пользователя
	userId, exist := ctx.Get("UserId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("UpdateAuthor: %v", "the user ID was not found in gin.Context"))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating the author.",
			nil,
		)
		return
	}
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("UpdateAuthor: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			nil,
		)
		return
	}
	// Обновляем данные о авторе
	errUpdateAuthor := as.repository.UpdateAuthor(body.Name, body.Description, id, int(userId.(int32)))

	if errUpdateAuthor != nil {
		logger.Log.Error(fmt.Sprintf("UpdateAuthor: %v", errUpdateAuthor))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating the author.",
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusCreated,
		"success",
		"OK.",
		nil,
	)
}
