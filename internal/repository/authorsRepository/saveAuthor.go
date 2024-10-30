package authorsRepository

import (
	"MySotre/pkg/pgDB"
)

func (ar *authorsRepository) SaveAuthor(
	name, description string,
	userId int,
	avatarId string,
) (int, error) {
	var authorId int

	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO authors(
			name, description, user_id, avatar_id
		) VALUES (
			$1, $2, $3, $4 
		) RETURNING id;
	`, name, description, userId, avatarId)

	if errQuery != nil {
		return -1, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&authorId); errScan != nil {
			return -1, errScan
		}
	}

	return authorId, nil
}
