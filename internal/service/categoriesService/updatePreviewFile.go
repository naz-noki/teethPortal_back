package categoriesService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update category preview file by file name
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category id"
// @Param file formData file true "New file"
// @Success 200 {file} file
// @Failure 400 {object} sendResponse.Response
// @Failure 404 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories/{id}/preview-file [put]
func (t *categoriesService) UpdatePreviewFile(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("UpdatePreviewFile: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid category id parameter.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Проверяем наличие записи по id
	category, errGetCategoryById := t.repository.GetCategoryById(id)

	if errGetCategoryById != nil {
		logger.Log.Error(fmt.Sprintf("UpdatePreviewFile: %v", errGetCategoryById.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error retrieving category information.",
			1,
			1,
			1,
			nil,
		)
		return
	} else if category.Id == 0 {
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"Category with this id not found.",
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
		logger.Log.Error(fmt.Sprintf("UpdatePreviewFile: %v", errFormFile.Error()))
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
	// Обновляем файл
	newPreviewFileId, errUpdatePreviewFile := t.repository.UpdatePreviewFile(category.PreviewFileId, fileHeader)

	if errUpdatePreviewFile != nil {
		logger.Log.Error(fmt.Sprintf("UpdatePreviewFile: %v", errUpdatePreviewFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error loading the new file.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Обновляем id файла для категории
	errUpdatePreviewFileId := t.repository.UpdatePreviewFileId(id, newPreviewFileId)

	if errUpdatePreviewFileId != nil {
		logger.Log.Error(fmt.Sprintf("UpdatePreviewFile: %v", errUpdatePreviewFileId.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error updating the category.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"OK.",
		1,
		1,
		1,
		nil,
	)
}
