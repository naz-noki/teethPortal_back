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

func CheckToken(tokensClient tokensApi.TokensClient) gin.HandlerFunc {
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
		checkTokenRequest := tokensApi.CheckTokenRequest{
			AccessToken: token[1],
		}

		checkTokenResponse, errCheckToken := tokensClient.CheckToken(context.Background(), &checkTokenRequest)

		if errCheckToken != nil {
			logger.Log.Error(fmt.Sprintf("ChechToken: %v", errCheckToken.Error()))
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
		if !checkTokenResponse.Result {
			sendResponse.Send(
				ctx,
				http.StatusUnauthorized,
				"error",
				"Invalid token.",
				nil,
			)
			return
		}

		ctx.Next()
	}
}
