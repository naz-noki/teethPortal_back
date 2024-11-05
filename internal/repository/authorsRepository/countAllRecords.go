package authorsRepository

import "MySotre/pkg/pgDB"

func (ar *authorsRepository) CountAllRecords() (int, error) {
	var counter int

	rows, errQuery := pgDB.DB.Query(`
		SELECT COUNT(id) 
		FROM authors; 
	`)

	if errQuery != nil {
		return -1, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&counter); errScan != nil {
			return -1, errScan
		}
	}

	return counter, nil
}
