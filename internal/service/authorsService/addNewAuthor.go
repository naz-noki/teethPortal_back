package authorsService

import (
	"MySotre/internal/config"
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *authorsService) AddNewAuthor(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.AddNewAuthorBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("AddNewAuthor: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Получаем id пользователя по его login
	userId, errGetUserIdByLogin := a.repository.GetUserIdByLogin(body.Login)

	if errGetUserIdByLogin != nil {
		logger.Log.Error(fmt.Sprintf("AddNewAuthor: %v", errGetUserIdByLogin.Error()))
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"An error occurred while receiving the user.",
			nil,
		)
		return
	}

	// Создаём аватар автора
	avatarPath, errSaveAuthorAvatar := a.repository.SaveAuthorAvatar(body.Name, config.Config.Data.AvatarsPath, body.AvatarName, body.AvatarData)

	if errSaveAuthorAvatar != nil {
		logger.Log.Error(fmt.Sprintf("AddNewAuthor: %v", errSaveAuthorAvatar.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the avatar for author.",
			nil,
		)
		return
	}

	// Сохраняем автора
	errSaveAuthor := a.repository.SaveAuthor(body.Name, body.Description, avatarPath, userId)

	if errSaveAuthor != nil {
		logger.Log.Error(fmt.Sprintf("AddNewAuthor: %v", errSaveAuthor.Error()))
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
		"OK",
		nil,
	)
}
