package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/paginationParams"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all authors
// @Tags authors
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} []service.GetAuthorByIdResponse
// @Failure 500 {object} sendResponse.Response
// @Router /api/authors/ [get]
func (as *authorsService) GetAllAuthors(ctx *gin.Context) {
	result := make([]*service.GetAuthorByIdResponse, 0, 9)
	// Получаем параметры пагинации
	pagination := paginationParams.Get(ctx)
	// Получаем все записи
	data, errGetAllAuthors := as.repository.GetAllAuthors(pagination.Limit, pagination.Offset)

	if errGetAllAuthors != nil {
		logger.Log.Error(fmt.Sprintf("GetAllAuthors: %v", errGetAllAuthors))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error getting the authors.",
			pagination.Page,
			pagination.Limit,
			0,
			nil,
		)
		return
	}

	for i := 0; i < len(data); i++ {
		result = append(
			result,
			&service.GetAuthorByIdResponse{
				AvatarId:    data[i].AvatarId,
				Description: data[i].Description,
				Id:          data[i].Id,
				Name:        data[i].Name,
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
		result,
	)
}
