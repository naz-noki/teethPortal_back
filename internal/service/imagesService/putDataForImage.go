package imagesService

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

func (i *imagesService) PutDataForImage(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.PutDataForImageBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("PutDataForImage: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Получаем imageId из параметров
	imageId, exist := ctx.Params.Get("imageId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("PutDataForImage: %v", errors.New("the image Id parameter was not found")))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid imageID parameter.",
			nil,
		)
		return
	}

	// Переводим imageId из string в int
	id, errAtoi := strconv.Atoi(imageId)

	if errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("PutDataForImage: %v", errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"Invalid imageID parameter.",
			nil,
		)
		return
	}

	// Обновляем данные о изображении в БД
	if errUpdateImageData := i.repository.UpdateImageData(id, body.Title, body.Description); errUpdateImageData != nil {
		logger.Log.Error(fmt.Sprintf("PutDataForImage: %v", errUpdateImageData.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while update data about image.",
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
