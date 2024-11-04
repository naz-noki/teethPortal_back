package authorsRepository

import "MySotre/pkg/pgDB"

func (ar *authorsRepository) CheckExistAuthor(id int) (bool, error) {
	idx := -1

	rows, errQuery := pgDB.DB.Query(`
		SELECT * FROM authors 
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
