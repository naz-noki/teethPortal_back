package middlewares

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func CheckIsAdmin(tokensClient tokensApi.TokensClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Получаем access токен из заголовка
		token := strings.Split(
			ctx.GetHeader("Authorization"),
			" ",
		)

		// Проверяем формат токена
		if len(token) != 2 || token[0] != "Bearer" {
			sendResponse.Send(
				ctx,
				http.StatusUnauthorized,
				"error",
				"Invalid token.",
				nil,
			)
			return
		}

		// Отправляем токен на проверку
		isAdminRequest := tokensApi.IsAdminRequest{
			AccessToken: token[1],
		}

		isAdminResponse, errIsAdmin := tokensClient.IsAdmin(context.Background(), &isAdminRequest)

		if errIsAdmin != nil {
			logger.Log.Error(fmt.Sprintf("CheckIsAdmin: %v", errIsAdmin.Error()))
			sendResponse.Send(
				ctx,
				http.StatusInternalServerError,
				"error",
				"An error occurred while verifying the token.",
				nil,
			)
			return
		}

		// Проверяем результат проверки токена
		if !isAdminResponse.Result {
			sendResponse.Send(
				ctx,
				http.StatusUnauthorized,
				"error",
				"The user is not an admin.",
				nil,
			)
			return
		}

		ctx.Next()
	}
}
