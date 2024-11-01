package artsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all arts
// @Tags arts
// @Accept json
// @Produce json
// @Success 200 {object} []service.GetArtResponse
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts [get]
func (as *artsService) GetArts(ctx *gin.Context) {
	resp := make([]*service.GetArtResponse, 0, 9)
	// Получаем все записи
	data, errGetArts := as.repository.GetArts()

	if errGetArts != nil {
		logger.Log.Error(fmt.Sprintf("GetArts: %v", errGetArts))
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
		fileIds, errGetFileIds := as.repository.GetFileIds(data[i].Id)

		if errGetFileIds != nil {
			logger.Log.Error(fmt.Sprintf("GetArts: %v", errGetFileIds))
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
