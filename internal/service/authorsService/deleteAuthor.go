package authorsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (as *authorsService) DeleteAuthor(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			nil,
		)
		return
	}
	// Получаем id аватара автора
	avatarId, errGetAvatarId := as.repository.GetAvatarId(id)

	if errGetAvatarId != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errGetAvatarId.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while deleting avatar for author.",
			nil,
		)
		return
	}
	// Удаляем аватар автора (файл)
	errDeleteAvatar := as.repository.DeleteAvatar(avatarId)

	if errDeleteAvatar != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errDeleteAvatar.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while deleting avatar for author.",
			nil,
		)
		return
	}
	// Удаляем автора
	errDeleteAuthor := as.repository.DeleteAuthor(id)

	if errDeleteAuthor != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errDeleteAuthor))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while deleting author.",
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusCreated,
		"error",
		"OK.",
		nil,
	)
}
