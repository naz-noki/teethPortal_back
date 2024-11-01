package ssoService

import (
	"MySotre/internal/repository/ssoRepository"
	"MySotre/internal/service"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

type ssoService struct {
	repository service.SsoRepository
	authApi.UnimplementedAuthServer
	tokensApi.UnimplementedTokensServer
}

func New() *ssoService {
	return &ssoService{
		repository: ssoRepository.New(),
	}
}
