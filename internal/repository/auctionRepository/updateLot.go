package auctionRepository

import (
	"MySotre/pkg/pgDB"
	"time"
)

func (ar *auctionRepository) UpdateLot(
	auctionId int,
	startTime, endTime time.Duration,
	price, artId, authorId, userId int,
	paid, sent bool,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE auction
		SET start_time = $1,end_time = $2,
		price = $3, lot = $4, seller = $5,
		buyer = $6, paid = $7, sent = $8
		WHERE id = $9;
	`, startTime, endTime, price, artId, authorId, userId, paid, sent, auctionId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
