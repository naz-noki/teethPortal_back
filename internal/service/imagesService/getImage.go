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

func (i *imagesService) GetImage(ctx *gin.Context) {
	// Получаем imageId из параметров
	imageId, exist := ctx.Params.Get("imageId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("GetImage: %v", errors.New("the image Id parameter was not found")))
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
		logger.Log.Error(fmt.Sprintf("GetImage: %v", errAtoi))
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
		logger.Log.Error(fmt.Sprintf("GetImage: %v", errGetImagePathById.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while receiving the image.",
			nil,
		)
		return
	}

	// Отправляем файл
	ctx.File(imagePath)
}
