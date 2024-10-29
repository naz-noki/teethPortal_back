package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (as *authorsService) SaveAuthor(ctx *gin.Context) {
	// Получаем id пользователя
	userId, exist := ctx.Get("UserId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("SaveAuthor: %v", "the user ID was not found in gin.Context"))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the author.",
			nil,
		)
		return
	}

	// Парсим тело запроса
	body := new(service.SaveAuthorBody)
	errUnmarshal := json.Unmarshal(
		[]byte(ctx.PostForm("data")),
		body,
	)

	if errUnmarshal != nil {
		logger.Log.Error(fmt.Sprintf("SaveAuthor: %v", errUnmarshal.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Получаем файл из запроса
	fileHeader, errFormFile := ctx.FormFile("file")

	if errFormFile != nil {
		logger.Log.Error(fmt.Sprintf("SaveAuthor: %v", errFormFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Сохраняем аватар автора
	avatarId, errSaveAvatar := as.repository.SaveAvatar(fileHeader)

	if errSaveAvatar != nil {
		logger.Log.Error(fmt.Sprintf("SaveAuthor: %v", errSaveAvatar))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the author avatar.",
			nil,
		)
		return
	}

	// Сохраняем автора
	_, errSaveAuthor := as.repository.SaveAuthor(body.Name, body.Description, int(userId.(int32)), avatarId)

	if errSaveAuthor != nil {
		logger.Log.Error(fmt.Sprintf("SaveAuthor: %v", errSaveAuthor))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the author.",
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
