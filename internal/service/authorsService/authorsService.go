package authorsService

import (
	"MySotre/internal/repository/authorsRepository"
	"MySotre/internal/service"
)

type authorsService struct {
	repository service.AuthorsRepository
}

func New() *authorsService {
	as := &authorsService{
		repository: authorsRepository.New(),
	}

	return as
}
