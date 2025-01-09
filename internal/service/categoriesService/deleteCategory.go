package categoriesService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Delete category by id
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories/{id} [delete]
func (t *categoriesService) DeleteCategory(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("DeleteCategory: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid id parameter.",
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
		logger.Log.Error(fmt.Sprintf("DeleteCategory: %v", errGetCategoryById.Error()))
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
	// Удаляем файл
	errDeletePreviewFile := t.repository.DeletePreviewFile(category.PreviewFileId)

	if errDeletePreviewFile != nil {
		logger.Log.Error(fmt.Sprintf("DeleteCategory: %v", errDeletePreviewFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while deleting the category preview file.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Удаляем запись
	errDeleteCategory := t.repository.DeleteCategory(category.Id)

	if errDeleteCategory != nil {
		logger.Log.Error(fmt.Sprintf("DeleteCategory: %v", errDeleteCategory.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while deleting the category.",
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
