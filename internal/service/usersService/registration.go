package usersService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
)

func (u *usersService) Registartion(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.RegistartionBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("Registartion: %v", errBindJSON.Error()))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}

	// Создаём запрос для GRPC сервера
	registrationRequest := authApi.RegistrationRequest{
		Login:    body.Login,
		Password: body.Password,
		IsAdmin:  body.IsAdmin,
	}
	// Делаем запрос на GRPC сервер для регистрации пользователя
	registrationResponse, errRegistration := u.client.Registration(context.Background(), &registrationRequest)

	if errRegistration != nil {
		logger.Log.Error(fmt.Sprintf("Registartion: %v", errRegistration.Error()))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred during user registration.",
			nil,
		)
		return
	}

	logger.Log.Info(fmt.Sprintf("Create user with id: %v", registrationResponse.GetUserId()))
	sendResponse.Send(
		ctx,
		http.StatusCreated,
		"success",
		"The user has been successfully created.",
		nil,
	)
}
