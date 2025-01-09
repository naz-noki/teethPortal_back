package categoriesService

import (
	"MySotre/internal/repository"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} []repository.Category
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories [get]
func (t *categoriesService) GetCategories(ctx *gin.Context) {
	// Получаем id всех категорий
	categoriesIds, errGetAllCategoriesIds := t.repository.GetAllCategoriesIds()

	if errGetAllCategoriesIds != nil {
		logger.Log.Error(fmt.Sprintf("GetCategories: %v", errGetAllCategoriesIds.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error retrieving all categories.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Получаем данные из БД о каждой категории
	categories := make([]*repository.Category, len(categoriesIds))

	for i := 0; i < len(categoriesIds); i++ {
		category, errGetCategoryById := t.repository.GetCategoryById(categoriesIds[i])

		if errGetCategoryById != nil {
			logger.Log.Error(fmt.Sprintf("GetCategories: %v", errGetCategoryById.Error()))
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
		}

		categories[i] = category
	}

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		1,
		1,
		1,
		categories,
	)
}
