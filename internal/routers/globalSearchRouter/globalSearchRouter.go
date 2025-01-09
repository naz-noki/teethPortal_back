package globalSearchRouter

import (
	"MySotre/internal/service/globalSearchService"

	"github.com/gin-gonic/gin"
)

func AddGlobalSearchRouter(s *gin.Engine) {
	gs := globalSearchService.New()

	globalSearch := s.Group("/api/global-search")
	{
		globalSearch.GET("/", gs.Search)
	}
}
