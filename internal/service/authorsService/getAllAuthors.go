package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *authorsService) GetAllAuthors(ctx *gin.Context) {
	resp := new(service.GetAllAuthorsResponse)
	// Получаем всех авторов из кеша

	// Получаем всех авторов
	authorsPointer, errGetAuthors := a.repository.GetAuthors()

	if errGetAuthors != nil {
		logger.Log.Error(fmt.Sprintf("GetAllAuthors: %v", errGetAuthors.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while receiving the authors.",
			nil,
		)
		return
	}

	// Переводим полученных авторов из БД в нужный вид для пользователя
	authors := *authorsPointer
	respAuthors := make([]service.AuthorForUser, len(authors))

	for i := 0; i < len(authors); i++ {
		author := service.AuthorForUser{
			Id:          authors[i].Id,
			Name:        authors[i].Name,
			Description: authors[i].Description,
		}

		respAuthors[i] = author
	}

	// Отправляем данные о всех авторах
	resp.Authors = &respAuthors
	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"OK",
		&resp,
	)
}
