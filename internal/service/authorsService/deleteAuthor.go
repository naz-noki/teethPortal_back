package authorsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *authorsService) DeleteAuthor(ctx *gin.Context) {
	// Получаем authorId из параметров
	authorId, exist := ctx.Params.Get("authorId")

	if !exist {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errors.New("the author Id parameter was not found")))
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
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errAtoi))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"Invalid authorId parameter.",
			nil,
		)
		return
	}

	// Получаем данные о авторе
	author, errGetAuthorById := a.repository.GetAuthorById(id)

	if errGetAuthorById != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errGetAuthorById))
		sendResponse.Send(
			ctx,
			http.StatusNotFound,
			"error",
			"An error occurred when deleting the author.",
			nil,
		)
		return
	}

	// Удаляем аватар автора
	errDeleteAuthorAvatar := a.repository.DeleteAuthorAvatar(author.Avatar)

	if errDeleteAuthorAvatar != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errDeleteAuthorAvatar))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred when deleting the avatar for author.",
			nil,
		)
		return
	}

	// Удаляем данные о авторе
	errDeleteAuthor := a.repository.DeleteAuthor(id)

	if errDeleteAuthor != nil {
		logger.Log.Error(fmt.Sprintf("DeleteAuthor: %v", errDeleteAuthor))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred when deleting the author.",
			nil,
		)
		return
	}

	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"OK",
		nil,
	)
}
