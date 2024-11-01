package authService

import (
	"MySotre/internal/repository/usersRepository"
	"MySotre/internal/service"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

type authService struct {
	authClient   authApi.AuthClient
	tokensClient tokensApi.TokensClient
	repository   service.AuthRepository
}

func New(
	authClient authApi.AuthClient,
	tokensClient tokensApi.TokensClient,
) *authService {
	return &authService{
		authClient:   authClient,
		tokensClient: tokensClient,
		repository:   usersRepository.New(),
	}
}
