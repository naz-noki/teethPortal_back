package artsRouter

import (
	"MySotre/internal/middlewares"
	"MySotre/internal/service/artsService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddArtsRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	as := artsService.New()

	arts := s.Group("/api/arts")
	{
		arts.POST("/", middlewares.CheckIsAdmin(tokensClient), as.SaveArt)
		arts.GET("/", as.GetArts)
		arts.GET("/:id", as.GetArtById)
		arts.GET("/:id/file/:fileName", as.GetArtFile)
		arts.PUT("/:id", middlewares.CheckIsAdmin(tokensClient), as.UpdateArt)
		arts.PUT("/:id/file/:fileName", middlewares.CheckIsAdmin(tokensClient), as.UpdateFile)
		arts.DELETE("/:id", as.DeleteArt)
		arts.DELETE("/:id/file/:fileName", as.DeleteFile)
	}
}
