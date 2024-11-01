package artsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Delete art record
// @Tags arts
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Art id"
// @Success 201 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts/{id} [delete]
func (as *artsService) DeleteArt(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("DeleteArt: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid art id parameter.",
			nil,
		)
		return
	}
	// Получаем id всех файлов из записи
	fileIds, errGetFileIds := as.repository.GetFileIds(id)

	if errGetFileIds != nil {
		logger.Log.Error(fmt.Sprintf("DeleteArt: %v", errGetFileIds))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while deleting the files from record.",
			nil,
		)
		return
	}
	// Удаляем все файлы
	for i := 0; i < len(fileIds); i++ {
		errDeleteFile := as.repository.DeleteFile(fileIds[i])

		if errDeleteFile != nil {
			logger.Log.Error(fmt.Sprintf("DeleteArt: %v", errDeleteFile))
			sendResponse.Send(
				ctx,
				http.StatusInternalServerError,
				"error",
				"An error occurred while deleting the file from record.",
				nil,
			)
			return
		}
	}
	// Удаляем запись
	errDeleteArt := as.repository.DeleteArt(id)

	if errDeleteArt != nil {
		logger.Log.Error(fmt.Sprintf("DeleteArt: %v", errDeleteArt))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while deleting the record.",
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusCreated,
		"success",
		"OK.",
		nil,
	)
}
