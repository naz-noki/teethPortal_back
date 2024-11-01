package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (ar *authorsRepository) GetAuthorById(id int) (*repository.Author, error) {
	author := new(repository.Author)

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, name, description, avatar_id, user_id
		FROM authors WHERE id = $1;
	`, id)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&author.Id, &author.Name, &author.Description, &author.AvatarId, &author.UserId); errScan != nil {
			return nil, errScan
		}
	}

	return author, nil
}
