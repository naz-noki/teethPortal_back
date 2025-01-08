package categoriesService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create new category
// @Tags categories
// @Accept json
// @Produce json
// @Param file formData file true "New file"
// @Param data formData service.CreateCategoryBody true "Category information"
// @Success 201 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories [post]
func (t *categoriesService) CreateCategory(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.CreateCategoryBody)
	errUnmarshal := json.Unmarshal(
		[]byte(ctx.PostForm("data")),
		body,
	)

	if errUnmarshal != nil {
		logger.Log.Error(fmt.Sprintf("CreateCategory: %v", errUnmarshal.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	// Получаем файл из запроса
	fileHeader, errFormFile := ctx.FormFile("file")

	if errFormFile != nil {
		logger.Log.Error(fmt.Sprintf("CreateCategory: %v", errFormFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	// Сохраняем файл
	fileId, errSavePreviewFile := t.repository.SavePreviewFile(fileHeader)

	if errSavePreviewFile != nil {
		logger.Log.Error(fmt.Sprintf("CreateCategory: %v", errSavePreviewFile))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the preview file.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	// Сохраняем автора
	errCreateCategory := t.repository.CreateCategory(body.Name, body.Description, fileId)

	if errCreateCategory != nil {
		logger.Log.Error(fmt.Sprintf("CreateCategory: %v", errCreateCategory))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving the data.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusCreated,
		"success",
		"OK.",
		1,
		1,
		1,
		nil,
	)
}
