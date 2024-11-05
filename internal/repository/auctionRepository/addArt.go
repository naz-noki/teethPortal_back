package auctionRepository

import "time"

func (ar *auctionRepository) AddArt(
	startTime, endTime time.Duration,
	price, artId, authorId, userId int,
	paid, sent bool,
) error {
	return nil
}
