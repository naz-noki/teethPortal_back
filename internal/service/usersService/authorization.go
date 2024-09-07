package usersService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
)

func (u *usersService) Authorization(ctx *gin.Context) {
	response := new(service.AuthorizationResponse)
	// Парсим тело запроса
	body := new(service.AuthorizationBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("Authorization: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Создаём запрос для GRPC сервера
	authorizationRequest := authApi.AuthorizationRequest{
		Login:    body.Login,
		Password: body.Password,
	}
	// Делаем запрос на GRPC сервер для авторизации пользователя
	authorizationResponse, errAuthorization := u.client.Authorization(context.Background(), &authorizationRequest)

	if errAuthorization != nil {
		logger.Log.Error(fmt.Sprintf("Authorization: %v", errAuthorization.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred during user authorization.",
			nil,
		)
		return
	}

	// Отправляем токены
	response.AccessToken = authorizationResponse.AccessToken
	// Устанавливаем значение refresh_token в куки пользователя
	ctx.SetCookie(
		"refreshToken",
		authorizationResponse.RefreshToken,
		3600*15, // 15 дней
		"/",
		"localhost",
		false,
		true, // http only
	)

	sendResponse.Send(
		ctx,
		http.StatusAccepted,
		"success",
		"The user has been successfully authorized.",
		response,
	)
}
