package imagesService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (i *imagesService) GetAllImages(ctx *gin.Context) {
	resp := new(service.GetAllImagesResponse)

	// Получаем данные о всех изображениях из БД
	imagesPointer, errGetImagesFromDB := i.repository.GetImagesFromDB()

	if errGetImagesFromDB != nil {
		logger.Log.Error(fmt.Sprintf("GetAllImages: %v", errGetImagesFromDB.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while receiving images.",
			nil,
		)
		return
	}

	// Записываем данные о изображениях в response
	images := *imagesPointer
	respImages := make([]service.ImageForUser, len(images))

	for i := 0; i < len(images); i++ {
		image := service.ImageForUser{
			Id:          images[i].Id,
			Title:       images[i].Title,
			Description: images[i].Description,
			CreatedAt:   images[i].CreatedAt,
		}

		respImages[i] = image
	}

	resp.Images = &respImages

	// Отправляем данные о всех изображениях
	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"OK",
		&resp,
	)
}
