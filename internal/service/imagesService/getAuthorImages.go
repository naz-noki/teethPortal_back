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

func (i *imagesService) GetAuthorImages(ctx *gin.Context) {
	resp := new(service.GetAllImagesResponse)

	// Получаем authorId из параметров
	authorId, exist := ctx.Params.Get("authorId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("GetAuthorImages: %v", errors.New("the author Id parameter was not found")))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid authorId parameter.",
			nil,
		)
		return
	}

	// Переводим authorId из string в int
	id, errAtoi := strconv.Atoi(authorId)

	if errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthorImages: %v", errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"Invalid authorId parameter.",
			nil,
		)
		return
	}

	// Получаем данные о всех изображениях автора из БД
	imagesPointer, errGetImagesFromDB := i.repository.GetImagesFromDBbyAuthorId(id)

	if errGetImagesFromDB != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthorImages: %v", errGetImagesFromDB.Error()))
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
