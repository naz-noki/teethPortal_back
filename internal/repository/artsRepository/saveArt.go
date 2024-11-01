package artsRepository

import (
	"MySotre/pkg/pgDB"
)

func (ar *artsRepository) SaveArt(
	title, description, content,
	authorId, artType string,
) error {
	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO arts (
			title, description, content, author_id, type
		) VALUES (
			$1, $2, $3, $4, $5 
		);
	`, title, description, content, authorId, artType)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
