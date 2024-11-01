package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all author arts
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} []service.GetArtResponse
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/authors/{id}/arts [get]
func (as *authorsService) GetAuthorArts(ctx *gin.Context) {
	resp := make([]*service.GetArtResponse, 0, 9)
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthorArts: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			nil,
		)
		return
	}
	// Получаем все записи
	data, errGetArts := as.artsRepository.GetAuthorArts(id)

	if errGetArts != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthorArts: %v", errGetArts))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting record.",
			nil,
		)
		return
	}
	// Получаем названия всех файлов для каждой записи
	for i := 0; i < len(data); i++ {
		fileIds, errGetFileIds := as.artsRepository.GetFileIds(data[i].Id)

		if errGetFileIds != nil {
			logger.Log.Error(fmt.Sprintf("GetAuthorArts: %v", errGetFileIds))
			sendResponse.Send(
				ctx,
				http.StatusInternalServerError,
				"error",
				"An error occurred while getting file idx for record.",
				nil,
			)
			return
		}

		resp = append(
			resp,
			&service.GetArtResponse{
				Id:          data[i].Id,
				Title:       data[i].Title,
				Description: data[i].Description,
				Content:     data[i].Content,
				AuthorId:    data[i].AuthorId,
				Type:        data[i].Type,
				Files:       fileIds,
			},
		)
	}

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		resp,
	)
}
