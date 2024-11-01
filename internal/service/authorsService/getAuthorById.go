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

// @Summary Get author by id
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} service.GetAuthorByIdResponse
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/authors/{id} [get]
func (as *authorsService) GetAuthorById(ctx *gin.Context) {
	resp := new(service.GetAuthorByIdResponse)

	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthorById: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			nil,
		)
		return
	}

	// Получаем автора
	author, errGetAuthorById := as.repository.GetAuthorById(id)

	if errGetAuthorById != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthorById: %v", errGetAuthorById))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error getting the author.",
			nil,
		)
		return
	}

	resp.AvatarId = author.AvatarId
	resp.Description = author.Description
	resp.Id = author.Id
	resp.Name = author.Name

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		resp,
	)
}
