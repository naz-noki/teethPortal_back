package artsService

import (
	"MySotre/internal/repository/artsRepository"
	"MySotre/internal/service"
)

type artsService struct {
	repository service.ArtsRepository
}

func New() *artsService {
	return &artsService{
		repository: artsRepository.New(),
	}
}
