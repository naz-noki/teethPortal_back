package imagesService

import (
	"MySotre/internal/config"
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (i *imagesService) AddNewImage(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.AddNewImageBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("AddNewImage: %v", errBindJSON.Error()))
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
	userId, errGetUserIdByLogin := i.repository.GetUserIdByLogin(body.Login)

	if errGetUserIdByLogin != nil {
		logger.Log.Error(fmt.Sprintf("AddNewImage: %v", errGetUserIdByLogin.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while receiving the user.",
			nil,
		)
		return
	}
	// Получаем id автора по его name
	authorId, errGetAuthorIdByName := i.repository.GetAuthorIdByName(body.AuthorName)

	if errGetAuthorIdByName != nil {
		logger.Log.Error(fmt.Sprintf("AddNewImage: %v", errGetAuthorIdByName.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while receiving the author.",
			nil,
		)
		return
	}

	// Сохраняем изображение
	filePath, errSaveImage := i.repository.SaveImage(config.Config.Data.ImagesPath, body.FileName, authorId, &body.File)

	if errSaveImage != nil {
		logger.Log.Error(fmt.Sprintf("AddNewImage: %v", errSaveImage.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the image.",
			nil,
		)
		return
	}

	// Сохраняем данные о изображении в БД
	if errSaveImageData := i.repository.SaveImageData(userId, authorId, filePath, body.Title, body.Description, body.CreatedAt); errSaveImageData != nil {
		logger.Log.Error(fmt.Sprintf("AddNewImage: %v", errSaveImageData.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the image.",
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
