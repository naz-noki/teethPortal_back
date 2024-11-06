package auctionRepository

import "MySotre/pkg/pgDB"

func (ar *auctionRepository) PlaceBet(
	auctionId, price, userId int,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE auction 
		SET price = $1, buyer = $2
		WHERE id = $3;
	`, price, userId, auctionId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
