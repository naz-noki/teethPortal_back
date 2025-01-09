package categoriesService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Update category by id
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body service.CreateCategoryBody true "New category information"
// @Success 200 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 404 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories/{id} [put]
func (t *categoriesService) UpdateCategory(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("UpdateCategory: Parameter id exist - %t, error: %v", existId, errAtoi))
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
	// Парсим тело запроса
	body := new(service.CreateCategoryBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("UpdateCategory: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Проверяем наличие категории
	existCategory, errCheckExistCategory := t.repository.CheckExistCategory(id)

	if errCheckExistCategory != nil {
		logger.Log.Error(fmt.Sprintf("UpdateCategory: %v", errCheckExistCategory.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while checking the existence of the category.",
			1,
			1,
			1,
			nil,
		)
		return
	} else if !existCategory {
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"Category with this ID not found.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Обновляем запись
	errUpdateCategory := t.repository.UpdateCategory(id, body.Name, body.Description)

	if errUpdateCategory != nil {
		logger.Log.Error(fmt.Sprintf("UpdateCategory: %v", errUpdateCategory.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating the category.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"OK.",
		1,
		1,
		1,
		nil,
	)
}
