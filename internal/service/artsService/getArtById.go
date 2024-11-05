package artsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get art by id
// @Tags arts
// @Accept json
// @Produce json
// @Param id path int true "Art ID"
// @Success 200 {object} service.GetArtResponse
// @Failure 400 {object} sendResponse.Response
// @Failure 404 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts/{id} [get]
func (as *artsService) GetArtById(ctx *gin.Context) {
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
			1,
			1,
			1,
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
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Получаем запись по id
	data, errGetArts := as.repository.GetArtById(id)

	if errGetArts != nil {
		logger.Log.Error(fmt.Sprintf("GetArtById: %v", errGetArts))
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"Art with this id - not found.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Получаем названия всех файлов для записи
	fileIds, errGetFileIds := as.repository.GetFileIds(data.Id)

	if errGetFileIds != nil {
		logger.Log.Error(fmt.Sprintf("GetArtById: %v", errGetFileIds))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting file idx for record.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	art := &service.GetArtResponse{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		AuthorId:    data.AuthorId,
		Type:        data.Type,
		Files:       fileIds,
	}

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		1,
		1,
		1,
		art,
	)
}
