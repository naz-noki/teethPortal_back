package categoriesService

import (
	"MySotre/internal/repository/categoriesRepository"
	"MySotre/internal/service"
)

type categoriesService struct {
	repository service.CategoriesRepository
}

func New() *categoriesService {
	return &categoriesService{
		repository: categoriesRepository.New(),
	}
}
