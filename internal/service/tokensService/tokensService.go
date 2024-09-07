package tokensService

import (
	"MySotre/internal/repository/tokensRepository"
	"MySotre/internal/service"

	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

type tokensService struct {
	client     tokensApi.TokensClient
	repository service.TokensRpository
}

func NewTokensService(tokensClient tokensApi.TokensClient) service.TokensService {
	ts := tokensService{
		client:     tokensClient,
		repository: tokensRepository.NewTokensRepository(),
	}

	return &ts
}
