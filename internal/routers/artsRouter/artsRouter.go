package artsRouter

import (
	"MySotre/internal/middlewares"
	"MySotre/internal/service/artsService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddArtsRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	as := artsService.New()

	tokens := s.Group("/api/arts")
	{
		tokens.POST("/", middlewares.CheckIsAdmin(tokensClient), as.SaveArt)
	}
}
