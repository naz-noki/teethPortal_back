package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *authorsService) PutAuthor(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.PutAuthorBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthor: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Получаем authorId из параметров
	authorId, exist := ctx.Params.Get("authorId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("PutAuthor: %v", errors.New("the author Id parameter was not found")))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid authorId parameter.",
			nil,
		)
		return
	}

	// Переводим authorId из string в int
	id, errAtoi := strconv.Atoi(authorId)

	if errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthor: %v", errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"Invalid authorId parameter.",
			nil,
		)
		return
	}

	// Обновляем данные о авторе
	errUpdateAuthor := a.repository.UpdateAuthor(body.Name, body.Description, id)

	if errUpdateAuthor != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthor: %v", errUpdateAuthor))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating information about the author.",
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"OK",
		nil,
	)
}
