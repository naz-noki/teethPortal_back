package auctionRepository

import (
	"MySotre/pkg/pgDB"
	"time"
)

func (ar *auctionRepository) AddLot(
	startTime, endTime time.Duration,
	price, artId, authorId int,
	paid, sent bool,
) error {
	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO auction (
			start_time,
			end_time,
			price,
			lot,
			seller,
			paid,
			sent
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7 
		);
	`, startTime, endTime, price, artId, authorId, paid, sent)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
