package imagesService

import (
	"MySotre/internal/repository/imagesRepository"
	"MySotre/internal/service"
)

type imagesService struct {
	repository service.ImagesRepository
}

func NewImagesService() service.ImagesService {
	is := imagesService{
		repository: imagesRepository.NewImagesRepository(),
	}

	return &is
}
