package artsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Delete art file
// @Tags arts
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Art id"
// @Param fileName path string true "File name"
// @Success 201 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts/{id}/file/{fileName} [delete]
func (as *artsService) DeleteFile(ctx *gin.Context) {
	// Получаем параметр fileName из запроса
	fileName, existFileName := ctx.Params.Get("fileName")

	if !existFileName {
		logger.Log.Error(fmt.Sprintf("DeleteFile: Parameter fileName exist - %t", existFileName))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid art fileName parameter.",
			nil,
		)
		return
	}
	// Удаляем файл
	errUpdateFile := as.repository.DeleteFile(fileName)

	if errUpdateFile != nil {
		logger.Log.Error(fmt.Sprintf("DeleteFile: %v", errUpdateFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while deleting file.",
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
