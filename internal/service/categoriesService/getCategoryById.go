package categoriesService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get category by id
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} repository.Category
// @Failure 400 {object} sendResponse.Response
// @Failure 404 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories/{id} [get]
func (t *categoriesService) GetCategoryById(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("GetCategoryById: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid id parameter.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Получаем данные из БД
	category, errGetCategoryById := t.repository.GetCategoryById(id)

	if errGetCategoryById != nil {
		logger.Log.Error(fmt.Sprintf("GetCategoryById: %v", errGetCategoryById.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error retrieving category information.",
			1,
			1,
			1,
			nil,
		)
		return
	} else if category.Id == 0 {
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"Category with this id not found.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		1,
		1,
		1,
		category,
	)
}
