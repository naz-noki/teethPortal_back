package authorsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *authorsService) GetAvatarForAuthor(ctx *gin.Context) {
	// Получаем authorId из параметров
	authorName, exist := ctx.Params.Get("authorName")

	if !exist {
		logger.Log.Error(fmt.Sprintf("GetAvatarForAuthor: %v", errors.New("the author Name parameter was not found")))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid authorName parameter.",
			nil,
		)
		return
	}

	// Получаем путь до изображения
	avatarPath, errGetImagePathById := a.repository.GetAvatarPathByAvtorName(authorName)

	if errGetImagePathById != nil {
		logger.Log.Error(fmt.Sprintf("GetImage: %v", errGetImagePathById.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while receiving the avatar.",
			nil,
		)
		return
	}

	// Отправляем файл
	ctx.File(avatarPath)
}
