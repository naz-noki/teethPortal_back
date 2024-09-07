package ssoService

import (
	"MySotre/internal/config"
	"MySotre/internal/service"
	"MySotre/pkg/cryptionPassword"
	"MySotre/pkg/jwtTokens"
	"context"
	"time"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ssoService) UpdateTokens(ctx context.Context, req *tokensApi.UpdateTokensRequest) (*tokensApi.UpdateTokensResponse, error) {
	resp := new(tokensApi.UpdateTokensResponse)
	user := new(service.UserPayload)
	// Получаем нужные параметры из UpdateTokensRequest
	id, refreshToken := req.GetUserId(), req.GetRefreshToken()
	// Получаем refresh токен из БД
	t, errGetRefreshTokenByUserId := s.repository.GetRefreshTokenByUserId(int(id))

	if errGetRefreshTokenByUserId != nil {
		return nil, status.Errorf(codes.Internal, "an error occurred while receiving a token from the database: %s", errGetRefreshTokenByUserId)
	}

	// Декодируем refresh токен из БД
	token, errDecode := cryptionPassword.Decode(t, config.Config.Tokens.RefreshSalt, config.Config.Tokens.RefreshSecondSalt)

	if errDecode != nil {
		return nil, status.Error(codes.Internal, "an error occurred while verifying tokens")
	}

	// Сравниваем токены
	if token != refreshToken {
		return nil, status.Error(codes.Unauthenticated, "invalid refresh token")
	}

	// Получаем данные пользователя для access токен по id
	u, errGetUserById := s.repository.GetUserById(int(id))

	if errGetUserById != nil {
		return nil, status.Error(codes.Internal, "error occurred while create the access token")
	}

	// Получаем access токен
	user.UserId = u.Id
	user.UserIsAdmin = u.IsAdmin
	accessToken, errCreateAccess := jwtTokens.CreateAccess(config.Config.Tokens.AccessSecret, &user, time.Minute*15)

	if errCreateAccess != nil {
		return nil, status.Error(codes.Internal, "error occurred while create the access token")
	}

	// Получаем refresh токен
	newRefreshToken := jwtTokens.CreateRefresh(config.Config.Tokens.RefreshSize)
	// Шифруем refresh токен
	rt := cryptionPassword.Encode(newRefreshToken, config.Config.Tokens.RefreshSalt, config.Config.Tokens.RefreshSecondSalt)
	// Сохраняем refresh токен в БД
	if errSetRefreshToken := s.repository.SetRefreshToken(int(id), rt); errSetRefreshToken != nil {
		return nil, status.Error(codes.Internal, "error occurred while create the refresh token")
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = newRefreshToken

	return resp, nil
}
