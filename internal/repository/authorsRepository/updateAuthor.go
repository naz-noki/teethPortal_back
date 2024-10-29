package authorsRepository

import (
	"MySotre/pkg/pgDB"
)

func (ar *authorsRepository) UpdateAuthor(
	name, description string,
	authorId, userId int,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE authors SET (
			name, description, user_id
		) VALUES (
			$1, $2, $3, 
		) WHERE id = $4  AND user_id = $3;
	`, name, description, userId, authorId)

	if errQuery != nil {
		rows.Close()
		return errQuery
	}
	defer rows.Close()

	return nil
}
