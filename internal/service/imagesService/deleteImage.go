package imagesService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (i *imagesService) DeleteImage(ctx *gin.Context) {
	// Получаем imageId из параметров
	imageId, exist := ctx.Params.Get("imageId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("DeleteImage: %v", errors.New("the image Id parameter was not found")))
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
		logger.Log.Error(fmt.Sprintf("DeleteImage: %v", errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"Invalid imageID parameter.",
			nil,
		)
		return
	}

	// Получаем путь до изображения
	imagePath, errGetImagePathById := i.repository.GetImagePathById(id)

	if errGetImagePathById != nil {
		logger.Log.Error(fmt.Sprintf("DeleteImage: %v", errGetImagePathById.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while delete the image.",
			nil,
		)
		return
	}

	// Удаляем изображение
	if errDeleteImage := i.repository.DeleteImage(imagePath); errDeleteImage != nil {
		logger.Log.Error(fmt.Sprintf("DeleteImage: %v", errDeleteImage.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while delete the image.",
			nil,
		)
		return
	}

	// Удаляем данные о изображении из БД
	if errDeleteImageData := i.repository.DeleteImageData(id, imagePath); errDeleteImageData != nil {
		logger.Log.Error(fmt.Sprintf("DeleteImage: %v", errDeleteImageData.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while delete the image.",
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
