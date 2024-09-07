package imagesRouter

import (
	"MySotre/internal/middlewares"
	"MySotre/internal/service/imagesService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddImagesRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	is := imagesService.NewImagesService()

	images := s.Group("/api/images")
	{
		images.POST("/add", middlewares.CheckIsAdmin(tokensClient), is.AddNewImage)
		images.GET("/all", is.GetAllImages)
		images.GET("/author/:authorId", is.GetAuthorImages)
		images.GET("/:imageId", is.GetImage)
		images.DELETE("/:imageId", middlewares.CheckIsAdmin(tokensClient), is.DeleteImage)
		images.PUT("/:imageId", middlewares.CheckIsAdmin(tokensClient), is.PutDataForImage)
	}
}
