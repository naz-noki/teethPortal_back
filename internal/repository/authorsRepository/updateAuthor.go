package authorsRepository

import (
	"MySotre/pkg/pgDB"
)

func (ar *authorsRepository) UpdateAuthor(
	name, description string,
	authorId, userId int,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE authors SET 
		name = $1, 
		description = $2, 
		user_id = $3
		WHERE id = $4 AND user_id = $3;
	`, name, description, userId, authorId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
