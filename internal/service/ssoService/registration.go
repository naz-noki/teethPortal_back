package ssoService

import (
	"MySotre/internal/config"
	"MySotre/pkg/cryptionPassword"
	"MySotre/pkg/logger"
	"context"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ssoService) Registration(ctx context.Context, req *authApi.RegistrationRequest) (*authApi.RegistrationResponse, error) {
	resp := new(authApi.RegistrationResponse)
	// Получаем нужные параметры из RegistrationRequest
	login, password, isAdmin := req.GetLogin(), req.GetPassword(), req.GetIsAdmin()
	// Проверяем логин
	if login == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login is empty")
	}
	// Проверяем пароль
	if password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password is empty")
	}

	// Шифруем пароль
	crptPass := cryptionPassword.Encode(password, config.Config.Password.Salt, config.Config.Password.SecondSalt)
	// Сохраняем пользователя в БД
	if errSetUser := s.repository.SetUser(login, crptPass, isAdmin); errSetUser != nil {
		logger.Log.Error(errSetUser.Error())
		return nil, status.Error(codes.AlreadyExists, "error when saving the user to the database")
	}

	// Получаем id пользователя (проверяем точно ли удалось сохранить пользователя)
	user, errGetUserIdByLogin := s.repository.GetUserByLogin(login)
	if errGetUserIdByLogin != nil {
		logger.Log.Error(errGetUserIdByLogin.Error())
		return nil, status.Errorf(codes.Internal, "error when check login: %v", errGetUserIdByLogin)
	}

	resp.UserId = int32(user.Id)
	return resp, nil
}
