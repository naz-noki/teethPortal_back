package authorsService

import (
	"MySotre/internal/repository/authorsRepository"
	"MySotre/internal/service"
)

type authorsService struct {
	repository service.AuthorsRepository
}

func NewAuthorsService() service.AuthorsService {
	as := authorsService{
		repository: authorsRepository.NewAuthorsRepository(),
	}

	return &as
}
