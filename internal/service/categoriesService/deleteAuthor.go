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

// @Summary Delete authors from category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body service.AddAuthorsBody true "Authors"
// @Success 201 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/categories/{id}/authors [delete]
func (t *categoriesService) DeleteAuthor(ctx *gin.Context) {
	// Получаем параметр id из запроса
	idParam, existId := ctx.Params.Get("id")
	id, errAtoi := strconv.Atoi(idParam)

	if !existId || errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: Parameter id exist - %t, error: %v", existId, errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author id parameter.",
			1,
			1,
			1,
			nil,
		)
		return
	}
	// Парсим тело запроса
	body := new(service.AddAuthorsBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errBindJSON.Error()))
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
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errCheckExistCategory.Error()))
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
	// Проверяем наличие авторов с такими ID
	for i := 0; i < len(body.AuthorsIds); i++ {
		authorId := body.AuthorsIds[i]
		existAuthor, errCheckExistAuthor := t.authorsRepository.CheckExistAuthor(authorId)

		if errCheckExistAuthor != nil {
			logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errCheckExistAuthor.Error()))
			sendResponse.Send(
				ctx,
				http.StatusInternalServerError,
				"error",
				"An error occurred while checking the existence of the author.",
				1,
				1,
				1,
				nil,
			)
			return
		} else if !existAuthor {
			sendResponse.Send(
				ctx,
				http.StatusNotFound,
				"error",
				"Author with this ID not found.",
				1,
				1,
				1,
				nil,
			)
			return
		}
		// Удаляем автора из категории
		if errDeleteAuthor := t.repository.DeleteAuthor(authorId, id); errDeleteAuthor != nil {
			logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errDeleteAuthor.Error()))
			sendResponse.Send(
				ctx,
				http.StatusInternalServerError,
				"error",
				"There was an error adding the author to the category.",
				1,
				1,
				1,
				nil,
			)
			return
		}
	}

	sendResponse.Send(
		ctx,
		http.StatusOK,
		"success",
		"OK.",
		1,
		1,
		1,
		nil,
	)
}
