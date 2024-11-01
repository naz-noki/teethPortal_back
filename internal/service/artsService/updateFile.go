package artsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update art file
// @Tags arts
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Art id"
// @Param fileName path string true "File name"
// @Param file formData file true "New file"
// @Success 200 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts/{id}/file/{fileName} [put]
func (as *artsService) UpdateFile(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("UpdateFile: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			nil,
		)
		return
	}
	// Получаем параметр fileName из запроса
	fileName, existFileName := ctx.Params.Get("fileName")

	if !existFileName {
		logger.Log.Error(fmt.Sprintf("UpdateFile: Parameter fileName exist - %t", existFileName))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author fileName parameter.",
			nil,
		)
		return
	}
	// Получаем файл из запроса
	fileHeader, errFormFile := ctx.FormFile("file")

	if errFormFile != nil {
		logger.Log.Error(fmt.Sprintf("UpdateFile: %v", errFormFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}
	// Обновляем файл
	errUpdateFile := as.repository.UpdateFile(id, fileName, fileHeader)

	if errUpdateFile != nil {
		logger.Log.Error(fmt.Sprintf("UpdateFile: %v", errUpdateFile.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while updating file for art.",
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		nil,
	)
}
