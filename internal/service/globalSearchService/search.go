package globalSearchService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Global search for all entities
// @Tags global search
// @Accept json
// @Produce json
// @Param query query string true "The string by which the search will take place"
// @Success 200 {object} service.SearchResponse
// @Failure 500 {object} sendResponse.Response
// @Router /api/global-search [get]
func (t *globalSearchService) Search(ctx *gin.Context) {
	query := ctx.Query("query")
	// Получаем список arts
	arts, errGetArts := t.repository.GetArts(query)

	if errGetArts != nil {
		logger.Log.Error(fmt.Sprintf("Search: %v", errGetArts.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error retrieving arts.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Получаем список authors
	authors, errGetAuthors := t.repository.GetAuthors(query)

	if errGetAuthors != nil {
		logger.Log.Error(fmt.Sprintf("Search: %v", errGetAuthors.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error retrieving authors.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Получаем список categories
	categories, errGetCategories := t.repository.GetCategories(query)

	if errGetCategories != nil {
		logger.Log.Error(fmt.Sprintf("Search: %v", errGetCategories.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"There was an error retrieving categories.",
			1,
			1,
			1,
			nil,
		)
		return
	}

	response := service.SearchResponse{
		Arts:       arts,
		Authors:    authors,
		Categories: categories,
	}
	sendResponse.Send(
		ctx,
		http.StatusInternalServerError,
		"success",
		"OK.",
		1,
		1,
		len(arts)+len(authors)+len(categories),
		&response,
	)
}
