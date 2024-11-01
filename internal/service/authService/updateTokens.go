package authService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

// @Summary Update user JWT - tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param body body service.UpdateTokensBody true "Access token"
// @Success 202 {object} service.UpdateTokensResponse
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/auth/tokens [post]
func (t *authService) UpdateTokens(ctx *gin.Context) {
	response := new(service.UpdateTokensResponse)
	// Парсим тело запроса
	body := new(service.UpdateTokensBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("UpdateTokens: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Получаем refresh token из cookie
	refreshToken, errCookie := ctx.Cookie("refreshToken")

	if errCookie != nil {
		logger.Log.Error(fmt.Sprintf("UpdateTokens: %v", errCookie.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Failed to get refresh token.",
			nil,
		)
		return
	}

	// Получаем id токена по login
	userId, errGetUserIdByLogin := t.repository.GetUserIdByLogin(body.Login)

	if errGetUserIdByLogin != nil {
		logger.Log.Error(fmt.Sprintf("UpdateTokens: %v", errGetUserIdByLogin.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating tokens.",
			nil,
		)
		return
	}

	// Создаём запрос для GRPC сервера
	updateTokensRequest := tokensApi.UpdateTokensRequest{
		UserId:       int32(userId),
		RefreshToken: refreshToken,
	}
	// Делаем запрос на GRPC сервер для авторизации пользователя
	updateTokensResponse, errUpdateTokens := t.tokensClient.UpdateTokens(context.Background(), &updateTokensRequest)

	if errUpdateTokens != nil {
		logger.Log.Error(fmt.Sprintf("UpdateTokens: %v", errUpdateTokens.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while updating tokens.",
			nil,
		)
		return
	}

	// Отправляем токены
	response.AccessToken = updateTokensResponse.AccessToken
	// Устанавливаем значение refresh_token в куки пользователя
	ctx.SetCookie(
		"refreshToken",
		updateTokensResponse.RefreshToken,
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
		"The tokens have been successfully updated.",
		response,
	)
}
