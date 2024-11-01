package authService

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

// @Summary Registration user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body service.RegistrationBody true "User information"
// @Success 201 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/auth/registration [post]
func (u *authService) Registration(ctx *gin.Context) {
	// Парсим тело запроса
	body := new(service.RegistrationBody)

	if errBindJSON := ctx.BindJSON(body); errBindJSON != nil {
		logger.Log.Error(fmt.Sprintf("Registration: %v", errBindJSON.Error()))
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
	registrationResponse, errRegistration := u.authClient.Registration(context.Background(), &registrationRequest)

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
