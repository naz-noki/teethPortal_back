package ssoService

import (
	"MySotre/internal/config"
	"MySotre/internal/service"
	"MySotre/pkg/jwtTokens"
	"context"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ssoService) IsAdmin(ctx context.Context, req *tokensApi.IsAdminRequest) (*tokensApi.IsAdminResponse, error) {
	resp := new(tokensApi.IsAdminResponse)
	userPayload := new(service.UserPayload)
	// Получаем нужные параметры из IsAdminRequest
	token := req.GetAccessToken()

	// Проверяем токен
	errCheckAccess := jwtTokens.CheckAccess(token, config.Config.Tokens.AccessSecret, userPayload)

	if errCheckAccess != nil {
		return nil, status.Errorf(codes.Unauthenticated, "error checking the token: %v", errCheckAccess)
	}

	// Проверяем является ли пользователь админом
	if userPayload.UserIsAdmin {
		resp.Result = true
	} else {
		resp.Result = false
	}

	resp.UserId = int32(userPayload.UserId)

	return resp, nil
}
