package authRouter

import (
	"MySotre/internal/service/authService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddAuthRoutes(
	s *gin.Engine,
	authClient authApi.AuthClient,
	tokensClient tokensApi.TokensClient,
) {
	authService := authService.New(authClient, tokensClient)

	auth := s.Group("/api/auth")
	{
		auth.POST("/registration", authService.Registration)
		auth.POST("/authorization", authService.Authorization)
		auth.POST("/tokens", authService.UpdateTokens)
	}
}
