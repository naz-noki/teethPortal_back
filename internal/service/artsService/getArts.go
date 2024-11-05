package artsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/paginationParams"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all arts
// @Tags arts
// @Accept json
// @Produce json
// @Param type query string false "Art type"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} []service.GetArtResponse
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts [get]
func (as *artsService) GetArts(ctx *gin.Context) {
	resp := make([]*service.GetArtResponse, 0, 9)
	// Получаем параметр типа записи
	artType := ctx.Query("type")
	// Получаем параметры пагинации
	pagination := paginationParams.Get(ctx)
	// Получаем все записи
	data, errGetArts := as.repository.GetArts(pagination.Limit, pagination.Offset, artType)

	if errGetArts != nil {
		logger.Log.Error(fmt.Sprintf("GetArts: %v", errGetArts))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting record.",
			pagination.Page,
			pagination.Limit,
			0,
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
				pagination.Page,
				pagination.Limit,
				0,
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
	// Получаем общее кол-во записей
	counter, errCountAllRecords := as.repository.CountAllRecords()

	if errCountAllRecords != nil {
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting counter for all records.",
			pagination.Page,
			pagination.Limit,
			0,
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		pagination.Page,
		pagination.Limit,
		counter,
		resp,
	)
}
