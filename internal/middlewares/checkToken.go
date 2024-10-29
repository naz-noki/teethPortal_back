package middlewares

import (
	"MySotre/pkg/logger"
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
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Отправляем токен на проверку
		checkTokenRequest := tokensApi.CheckTokenRequest{
			AccessToken: token[1],
		}

		checkTokenResponse, errCheckToken := tokensClient.CheckToken(context.Background(), &checkTokenRequest)
		if errCheckToken != nil {
			logger.Log.Error(fmt.Sprintf("ChechToken: %v", errCheckToken.Error()))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Проверяем результат проверки токена
		if !checkTokenResponse.Result {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("UserId", checkTokenResponse.UserId)
		ctx.Next()
	}
}
