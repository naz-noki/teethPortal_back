package authorsRouter

import (
	"MySotre/internal/middlewares"
	"MySotre/internal/service/authorsService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddAuthorsRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	as := authorsService.New()

	authors := s.Group("/api/authors")
	{
		authors.POST("/", middlewares.CheckIsAdmin(tokensClient), as.SaveAuthor)
		authors.GET("/", as.GetAllAuthors)
		authors.GET("/:id", as.GetAuthorById)
		authors.GET("/:id/arts", as.GetAuthorArts)
		authors.GET("/:id/avatar/:fileName", as.GetAvatar)
		authors.PUT("/:id", middlewares.CheckIsAdmin(tokensClient), as.UpdateAuthor)
		authors.PUT("/:id/avatar/:fileName", middlewares.CheckIsAdmin(tokensClient), as.UpdateAvatar)
		authors.DELETE("/:id", middlewares.CheckIsAdmin(tokensClient), as.DeleteAuthor)
	}
}
