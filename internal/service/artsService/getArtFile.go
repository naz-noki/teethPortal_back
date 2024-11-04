package artsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get art file by file name
// @Tags arts
// @Accept json
// @Produce application/octet-stream
// @Param id path int true "Art id"
// @Param fileName path string true "File name"
// @Success 200 {file} file
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts/{id}/file/{fileName} [get]
func (as *artsService) GetArtFile(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("GetArtById: Parameter id exist - %t, error: %v", existId, errAtoi))
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
		logger.Log.Error(fmt.Sprintf("GetArtFile: Parameter fileName exist - %t", existFileName))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author fileName parameter.",
			nil,
		)
		return
	}
	// Проверяем наличие записи по id
	check, errCheckExistArt := as.repository.CheckExistArt(id)

	if errCheckExistArt != nil || !check {
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"Art with this id - not found.",
			nil,
		)
		return
	}
	// Получаем файл
	val, errGetAvatar := as.repository.GetFile(fileName)

	if errGetAvatar != nil {
		logger.Log.Error(fmt.Sprintf("GetArtFile: %v", errGetAvatar))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting the author avatar.",
			nil,
		)
		return
	}

	stat, errStat := val.Stat()
	if errStat != nil {
		logger.Log.Error(fmt.Sprintf("GetArtFile: %v", errStat))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting the author avatar.",
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
