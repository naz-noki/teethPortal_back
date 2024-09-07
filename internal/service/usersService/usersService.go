package usersService

import (
	"MySotre/internal/service"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
)

type usersService struct {
	client authApi.AuthClient
}

func NewUsersService(authClient authApi.AuthClient) service.UsersService {
	us := usersService{
		client: authClient,
	}

	return &us
}
