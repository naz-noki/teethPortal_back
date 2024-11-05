package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/paginationParams"
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
// @Failure 404 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/authors/{id}/arts [get]
func (as *authorsService) GetAuthorArts(ctx *gin.Context) {
	resp := make([]*service.GetArtResponse, 0, 9)
	// Получаем параметры пагинации
	pagination := paginationParams.Get(ctx)
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
			pagination.Page,
			pagination.Limit,
			0,
			nil,
		)
		return
	}
	// Проверяем наличие записи по id
	check, errCheckExistAuthor := as.repository.CheckExistAuthor(id)

	if errCheckExistAuthor != nil || !check {
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"Author with this id - not found.",
			pagination.Page,
			pagination.Limit,
			0,
			nil,
		)
		return
	}
	// Получаем все записи
	data, errGetArts := as.artsRepository.GetAuthorArts(id, pagination.Limit, pagination.Offset)

	if errGetArts != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthorArts: %v", errGetArts))
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
		fileIds, errGetFileIds := as.artsRepository.GetFileIds(data[i].Id)

		if errGetFileIds != nil {
			logger.Log.Error(fmt.Sprintf("GetAuthorArts: %v", errGetFileIds))
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
