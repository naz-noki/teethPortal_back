package authorsService

import (
	"MySotre/internal/repository/artsRepository"
	"MySotre/internal/repository/authorsRepository"
	"MySotre/internal/service"
)

type authorsService struct {
	repository     service.AuthorsRepository
	artsRepository service.ArtsRepository
}

func New() *authorsService {
	as := &authorsService{
		repository:     authorsRepository.New(),
		artsRepository: artsRepository.New(),
	}

	return as
}
