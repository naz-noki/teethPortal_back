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
		tokens.GET("/", as.GetArts)
		tokens.GET("/:id", as.GetArtById)
		tokens.GET("/:id/file/:fileName", as.GetArtFile)
		tokens.PUT("/:id", middlewares.CheckIsAdmin(tokensClient), as.UpdateArt)
		tokens.PUT("/:id/file/:fileName", middlewares.CheckIsAdmin(tokensClient), as.UpdateFile)
		tokens.DELETE("/:id", as.DeleteArt)
		tokens.DELETE("/:id/file/:fileName", as.DeleteFile)
	}
}
