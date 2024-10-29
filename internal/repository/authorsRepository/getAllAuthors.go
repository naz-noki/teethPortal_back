package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (ar *authorsRepository) GetAllAuthors() ([]*repository.Author, error) {
	result := make([]*repository.Author, 0, 9)

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, name, description, avatar_id, user_id
		FROM authors;
	`)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		author := new(repository.Author)

		if errScan := rows.Scan(&author.Id, &author.Name, &author.Description, &author.AvatarId, &author.UserId); errScan != nil {
			return nil, errScan
		}

		result = append(result, author)
	}

	return result, nil
}
