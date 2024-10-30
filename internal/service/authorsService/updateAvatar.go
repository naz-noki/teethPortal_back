package authorsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (as *authorsService) UpdateAvatar(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("UpdateAvatar: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			nil,
		)
		return
	}
	// Получаем параметр fileName из запроса
	fileName, existFileName := ctx.Params.Get("fileName")

	if !existFileName {
		logger.Log.Error(fmt.Sprintf("UpdateAvatar: Parameter fileName exist - %t", existFileName))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author fileName parameter.",
			nil,
		)
		return
	}
	// Получаем файл из запроса
	fileHeader, errFormFile := ctx.FormFile("file")

	if errFormFile != nil {
		logger.Log.Error(fmt.Sprintf("UpdateAvatar: %v", errFormFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}
	// Обновляем аватар (файл)
	newFileName, errUpdateAvatar := as.repository.UpdateAvatar(fileName, fileHeader)

	if errUpdateAvatar != nil {
		logger.Log.Error(fmt.Sprintf("UpdateAvatar: %v", errUpdateAvatar.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while updateing avatar for author.",
			nil,
		)
		return
	}
	// Обновляем avatar_id для автора
	errUpdateAvatarId := as.repository.UpdateAvatarId(id, newFileName)

	if errUpdateAvatarId != nil {
		logger.Log.Error(fmt.Sprintf("UpdateAvatar: %v", errUpdateAvatarId.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while updateing avator for author.",
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
