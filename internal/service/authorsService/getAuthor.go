package authorsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *authorsService) GetAuthor(ctx *gin.Context) {
	resp := new(service.GetAuthorResponse)
	// Получаем authorId из параметров
	authorId, exist := ctx.Params.Get("authorId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("GetAuthor: %v", errors.New("the author Id parameter was not found")))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid authorId parameter.",
			nil,
		)
		return
	}

	// Переводим authorId из string в int
	id, errAtoi := strconv.Atoi(authorId)

	if errAtoi != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthor: %v", errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"Invalid authorId parameter.",
			nil,
		)
		return
	}

	// Получаем автора по id
	authorPointer, errGetAuthorById := a.repository.GetAuthorById(id)

	if errGetAuthorById != nil {
		logger.Log.Error(fmt.Sprintf("GetAuthor: %v", errGetAuthorById))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting information about the author.",
			nil,
		)
		return
	}

	// Отправляем результат
	author := service.AuthorForUser{
		Id:          authorPointer.Id,
		Name:        authorPointer.Name,
		Description: authorPointer.Description,
	}
	resp.Author = &author

	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"OK",
		resp,
	)
}
