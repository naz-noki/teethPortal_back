package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func CheckIsAdmin(tokensClient tokensApi.TokensClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// // Получаем access токен из заголовка
		// token := strings.Split(
		// 	ctx.GetHeader("Authorization"),
		// 	" ",
		// )

		// // Проверяем формат токена
		// if len(token) != 2 || token[0] != "Bearer" {
		// 	ctx.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// // Отправляем токен на проверку
		// isAdminRequest := tokensApi.IsAdminRequest{
		// 	AccessToken: token[1],
		// }

		// isAdminResponse, errIsAdmin := tokensClient.IsAdmin(context.Background(), &isAdminRequest)

		// if errIsAdmin != nil {
		// 	logger.Log.Error(fmt.Sprintf("CheckIsAdmin: %v", errIsAdmin.Error()))
		// 	ctx.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// // Проверяем результат проверки токена
		// if !isAdminResponse.Result {
		// 	ctx.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }

		// ctx.Set("UserId", isAdminResponse.UserId)
		// ctx.Next()

		ctx.Set("UserId", int32(2))
		ctx.Next()
	}
}
