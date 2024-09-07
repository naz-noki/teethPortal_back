package imagesRepository

import "MySotre/pkg/pgDB"

func (i *imagesRepository) GetAuthorIdByName(name string) (int, error) {
	id := -1

	rows, errQuery := pgDB.DB.Query(`
		SELECT id FROM authors 
		WHERE name = $1
	`, name)

	if errQuery != nil {
		return id, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		errScan := rows.Scan(&id)

		if errScan != nil {
			return id, errScan
		}
	}

	return id, nil
}
