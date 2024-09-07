package authorsRepository

import (
	"MySotre/pkg/pgDB"
)

func (a *authorsRepository) SaveAuthor(
	name, description, avatarPath string,
	userId int,
) error {
	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO authors (
			name,
			description,
			avatar,
			user_id
		) VALUES ($1, $2, $3, $4)
	`, name, description, avatarPath, userId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
