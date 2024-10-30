package authorsRouter

import (
	"MySotre/internal/middlewares"
	"MySotre/internal/service/authorsService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddAuthorsRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	as := authorsService.New()

	tokens := s.Group("/api/authors")
	{
		tokens.POST("/", middlewares.CheckIsAdmin(tokensClient), as.SaveAuthor)
		tokens.GET("/", as.GetAllAuthors)
		tokens.GET("/:id", as.GetAuthorById)
		tokens.GET("/:id/avatar/:fileName", as.GetAvatar)
		tokens.PUT("/:id", middlewares.CheckIsAdmin(tokensClient), as.UpdateAuthor)
		tokens.PUT("/:id/avatar/:fileName", middlewares.CheckIsAdmin(tokensClient), as.UpdateAvatar)
		tokens.DELETE("/:id", middlewares.CheckIsAdmin(tokensClient), as.DeleteAuthor)
	}
}
