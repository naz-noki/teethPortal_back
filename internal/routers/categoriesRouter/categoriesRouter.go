package categoriesRouter

import (
	"MySotre/internal/middlewares"
	"MySotre/internal/service/categoriesService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddCategoriesRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	cs := categoriesService.New()

	categories := s.Group("/api/categories")
	{
		categories.POST("/", middlewares.CheckIsAdmin(tokensClient), cs.CreateCategory)
		// categories.GET("/", as.GetAllAuthors)
		// categories.GET("/:id", as.GetAuthorById)
		// categories.GET("/:id/arts", as.GetAuthorArts)
		// categories.GET("/:id/avatar/:fileName", as.GetAvatar)
		// categories.PUT("/:id", middlewares.CheckIsAdmin(tokensClient), as.UpdateAuthor)
		// categories.PUT("/:id/avatar/:fileName", middlewares.CheckIsAdmin(tokensClient), as.UpdateAvatar)
		// categories.DELETE("/:id", middlewares.CheckIsAdmin(tokensClient), as.DeleteAuthor)
	}
}
