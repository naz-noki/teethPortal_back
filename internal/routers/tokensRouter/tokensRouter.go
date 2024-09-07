package tokensRouter

import (
	"MySotre/internal/service/tokensService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func AddTokensRoutes(s *gin.Engine, tokensClient tokensApi.TokensClient) {
	ts := tokensService.NewTokensService(tokensClient)

	tokens := s.Group("/api/tokens")
	{
		tokens.POST("/update", ts.UpdateTokens)
	}
}
