package ssoService

import (
	"MySotre/internal/config"
	"MySotre/internal/service"
	"MySotre/pkg/cryptionPassword"
	"MySotre/pkg/jwtTokens"
	"context"
	"time"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ssoService) Authorization(ctx context.Context, req *authApi.AuthorizationRequest) (*authApi.AuthorizationResponse, error) {
	resp := new(authApi.AuthorizationResponse)
	// Получаем нужные параметры из AuthorizationRequest
	login, password := req.GetLogin(), req.GetPassword()
	// Проверяем логин
	if login == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login is empty")
	}
	// Проверяем пароль
	if password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is empty")
	}

	// Получаем пользователя по login
	user, errGetUserByLogin := s.repository.GetUserByLogin(login)

	if errGetUserByLogin != nil {
		return nil, status.Error(codes.NotFound, "user with this login was not found")
	}

	// Декодируем пароль полученный из БД
	pass, errDecode := cryptionPassword.Decode(user.Password, config.Config.Password.Salt, config.Config.Password.SecondSalt)

	if errDecode != nil {
		return nil, status.Error(codes.Internal, "error occurred while verifying the password")
	}

	// Сравниваем пароли
	if pass != password {
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	// Получаем access токен
	accessTokenPayload := service.UserPayload{
		UserId:      user.Id,
		UserIsAdmin: user.IsAdmin,
	}
	accessToken, errCreateAccess := jwtTokens.CreateAccess(config.Config.Tokens.AccessSecret, &accessTokenPayload, time.Minute*15)

	if errCreateAccess != nil {
		return nil, status.Error(codes.Internal, "error occurred while create the access token")
	}

	// Получаем refresh токен
	refreshToken := jwtTokens.CreateRefresh(config.Config.Tokens.RefreshSize)
	// Шифруем refresh токен
	rt := cryptionPassword.Encode(refreshToken, config.Config.Tokens.RefreshSalt, config.Config.Tokens.RefreshSecondSalt)
	// Сохраняем refresh токен в БД
	if errSetRefreshToken := s.repository.SetRefreshToken(user.Id, rt); errSetRefreshToken != nil {
		return nil, status.Error(codes.Internal, "error occurred while create the refresh token")
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	return resp, nil
}
