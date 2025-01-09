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
		categories.GET("/", cs.GetCategories)
		categories.GET("/:id", cs.GetCategoryById)
		categories.GET("/:id/preview-file/:fileName", cs.GetPreviewFile)
		categories.POST("/", middlewares.CheckIsAdmin(tokensClient), cs.CreateCategory)
		categories.POST("/:id/authors", middlewares.CheckIsAdmin(tokensClient), cs.CreateCategory)
		categories.DELETE("/:id/authors", middlewares.CheckIsAdmin(tokensClient), cs.CreateCategory)
	}
}
