package categoriesService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get category preview file by file name
// @Tags categories
// @Accept json
// @Produce application/octet-stream
// @Param id path int true "Category id"
// @Param fileName path string true "File name"
// @Success 200 {file} file
// @Failure 400 {object} sendResponse.Response
// @Failure 404 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories/{id}/preview-file/{fileName} [get]
func (t *categoriesService) GetPreviewFile(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("GetPreviewFile: Parameter id exist - %t, error: %v", existId, errAtoi))
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
	// Получаем параметр fileName из запроса
	fileName, existFileName := ctx.Params.Get("fileName")

	if !existFileName {
		logger.Log.Error(fmt.Sprintf("GetPreviewFile: Parameter fileName exist - %t", existFileName))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid category fileName parameter.",
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
		logger.Log.Error(fmt.Sprintf("GetCategoryById: %v", errGetCategoryById.Error()))
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
	// Получаем файл
	val, errGetPreviewFile := t.repository.GetPreviewFile(fileName)

	if errGetPreviewFile != nil {
		logger.Log.Error(fmt.Sprintf("GetCategoryById: %v", errGetPreviewFile))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting preview file.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	stat, errStat := val.Stat()
	if errStat != nil {
		logger.Log.Error(fmt.Sprintf("GetCategoryById: %v", errStat))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting the preview file.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	defer val.Close()

	ctx.Status(http.StatusOK)
	ctx.Header("Content-Disposition", "attachment; filename="+stat.Key)
	ctx.Stream(func(w io.Writer) bool {
		_, errCopy := io.Copy(w, val)
		return errCopy == nil && false
	})
}
