package artsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update data for art
// @Tags arts
// @Accept json
// @Produce json
// @Param id path int true "Art id"
// @Param body body service.SaveArtBody true "New art information"
// @Success 200 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts/{id} [put]
func (as *artsService) UpdateArt(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.SaveArtBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("UpdateArt: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		log.Println("qwe")
		logger.Log.Error(fmt.Sprintf("UpdateArt: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			nil,
		)
		return
	}
	// Обновляем данные
	errUpdateArt := as.repository.UpdateArt(body.Title, body.Description, body.Content, body.Type, id, body.AuthorId)

	if errUpdateArt != nil {
		logger.Log.Error(fmt.Sprintf("UpdateArt: %v", errUpdateArt))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating the record.",
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
