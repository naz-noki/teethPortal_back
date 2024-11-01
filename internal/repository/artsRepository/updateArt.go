package artsRepository

import "MySotre/pkg/pgDB"

func (ar *artsRepository) UpdateArt(
	title, description,
	content, artType string,
	artId, authorId int,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE arts 
		SET title = $1, description = $2, content = $3, type = $4, author_id = $6
		WHERE id = $5 AND author_id = $6;
	`, title, description, content, artType, artId, authorId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
