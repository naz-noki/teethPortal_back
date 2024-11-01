package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all authors
// @Tags authors
// @Accept json
// @Produce json
// @Success 200 {object} []service.GetAuthorByIdResponse
// @Failure 500 {object} sendResponse.Response
// @Router /api/authors/ [get]
func (as *authorsService) GetAllAuthors(ctx *gin.Context) {
	result := make([]*service.GetAuthorByIdResponse, 0, 9)

	data, errGetAllAuthors := as.repository.GetAllAuthors()

	if errGetAllAuthors != nil {
		logger.Log.Error(fmt.Sprintf("GetAllAuthors: %v", errGetAllAuthors))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error getting the authors.",
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

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		result,
	)
}
