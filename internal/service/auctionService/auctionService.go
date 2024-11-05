package auctionService

import (
	"MySotre/internal/repository/auctionRepository"
	"MySotre/internal/service"
)

type auctionService struct {
	repository service.AuctionRepository
}

func New() *auctionService {
	return &auctionService{
		repository: auctionRepository.New(),
	}
}
