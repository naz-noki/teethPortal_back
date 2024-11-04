package artsRepository

import (
	"MySotre/pkg/pgDB"
	"log"
)

func (ar *artsRepository) CheckExistArt(id int) (bool, error) {
	log.Println("qwe")
	idx := -1

	rows, errQuery := pgDB.DB.Query(`
		SELECT id FROM arts 
		WHERE id = $1;
	`, id)

	if errQuery != nil {
		return false, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&idx); errScan != nil {
			return false, errScan
		}
	}

	if idx != -1 {
		return true, nil
	}
	return false, nil
}
