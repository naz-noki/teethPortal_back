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

func (a *authorsService) PutAuthorAvatar(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.PutAuthorAvatarBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthorAvatar: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Получаем путь к аватару по имени автора
	avatarPath, errGetAvatarPathByAvtorName := a.repository.GetAvatarPathByAvtorName(body.Name)

	if errGetAvatarPathByAvtorName != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthorAvatar: %v", errGetAvatarPathByAvtorName.Error()))
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"An error occurred while getting the author's avatar.",
			nil,
		)
		return
	}

	// Удаляем предыдущий аватар автора
	errDeleteAuthorAvatar := a.repository.DeleteAuthorAvatar(avatarPath)

	if errDeleteAuthorAvatar != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthorAvatar: %v", errDeleteAuthorAvatar.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating the author's avatar.",
			nil,
		)
		return
	}

	// Сохраняем новый аватар автора
	newAvatarPath, errSaveAuthorAvatar := a.repository.SaveAuthorAvatar(body.Name, config.Config.Data.AvatarsPath, body.AvatarName, body.AvatarData)

	if errSaveAuthorAvatar != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthorAvatar: %v", errSaveAuthorAvatar.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating the author's avatar.",
			nil,
		)
		return
	}

	// Обновляем путь до автара автора в БД
	errUpdatePathToAuthorAvatar := a.repository.UpdatePathToAuthorAvatar(newAvatarPath, body.Name)

	if errUpdatePathToAuthorAvatar != nil {
		logger.Log.Error(fmt.Sprintf("PutAuthorAvatar: %v", errUpdatePathToAuthorAvatar.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating the author's avatar.",
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusCreated,
		"error",
		"OK",
		nil,
	)
}
