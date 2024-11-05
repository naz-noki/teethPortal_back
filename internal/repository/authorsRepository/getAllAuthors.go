package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (ar *authorsRepository) GetAllAuthors(limit, offset int) ([]*repository.Author, error) {
	result := make([]*repository.Author, 0, 9)

	if limit == 0 {
		limit = repository.MaxLimit
	}

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, name, description, avatar_id, user_id
		FROM authors
		ORDER BY id LIMIT $1 OFFSET $2;
	`, limit, offset)

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
