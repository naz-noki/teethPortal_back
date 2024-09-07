package authorsRouter

import (
	"MySotre/internal/middlewares"
	"MySotre/internal/service/authorsService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddAuthorsRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	as := authorsService.NewAuthorsService()

	authors := s.Group("/api/authors")
	{
		authors.POST("/add", middlewares.CheckIsAdmin(tokensClient), as.AddNewAuthor)
		authors.GET("/all", as.GetAllAuthors)
		authors.GET("/:authorId", as.GetAuthor)
		authors.GET("/avatar/:authorName", as.GetAvatarForAuthor)
		authors.PUT("/:authorId", middlewares.CheckIsAdmin(tokensClient), as.PutAuthor)
		authors.PUT("/avatar", middlewares.CheckIsAdmin(tokensClient), as.PutAuthorAvatar)
		authors.DELETE("/:authorId", middlewares.CheckIsAdmin(tokensClient), as.DeleteAuthor)
	}
}
