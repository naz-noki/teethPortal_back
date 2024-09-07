package ssoService

import (
	"MySotre/internal/config"
	"MySotre/internal/service"
	"MySotre/pkg/jwtTokens"
	"context"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

func (s *ssoService) CheckToken(ctx context.Context, req *tokensApi.CheckTokenRequest) (*tokensApi.CheckTokenResponse, error) {
	resp := new(tokensApi.CheckTokenResponse)
	userPayload := new(service.UserPayload)
	// Получаем нужные параметры из CheckTokenRequest
	token := req.GetAccessToken()

	// Проверяем токен
	errCheckAccess := jwtTokens.CheckAccess(token, config.Config.Tokens.AccessSecret, userPayload)

	if errCheckAccess != nil {
		resp.Result = false
	} else {
		resp.Result = true
	}

	return resp, nil
}
