package globalSearchService

import (
	"MySotre/internal/repository/globalSearchRepository"
	"MySotre/internal/service"
)

type globalSearchService struct {
	repository service.GlobalSearchRepository
}

func New() *globalSearchService {
	return &globalSearchService{
		repository: globalSearchRepository.New(),
	}
}
