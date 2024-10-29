package usersService

import (
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
)

type usersService struct {
	client authApi.AuthClient
}

func NewUsersService(authClient authApi.AuthClient) *usersService {
	us := usersService{
		client: authClient,
	}

	return &us
}
