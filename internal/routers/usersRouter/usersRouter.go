package usersRouter

import (
	"MySotre/internal/service/usersService"

	"github.com/gin-gonic/gin"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
)

func AddUsersRoutes(s *gin.Engine, authClient authApi.AuthClient) {
	usersService := usersService.NewUsersService(authClient)

	users := s.Group("/api/users")
	{
		users.POST("/registration", usersService.Registartion)
		users.POST("/authorization", usersService.Authorization)
	}
}
