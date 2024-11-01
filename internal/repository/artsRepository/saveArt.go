package artsRepository

import (
	"MySotre/pkg/pgDB"
)

func (ar *artsRepository) SaveArt(
	title, description,
	content, artType string,
	authorId int,
) (int, error) {
	var artId int

	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO arts (
			title, description, content, author_id, type
		) VALUES (
			$1, $2, $3, $4, $5 
		) RETURNING id;
	`, title, description, content, authorId, artType)

	if errQuery != nil {
		return -1, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&artId); errScan != nil {
			return -1, errScan
		}
	}

	return artId, nil
}
