package artsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (ar *artsRepository) GetArts(
	limit, offset int,
	artType string,
) ([]*repository.Art, error) {
	result := make([]*repository.Art, 0, 9)

	if limit == 0 {
		limit = repository.MaxLimit
	}

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, title, description, content, author_id, type 
		FROM arts
		WHERE ($3 = '' OR type::text = $3) 
		ORDER BY id LIMIT $1 OFFSET $2; 
	`, limit, offset, artType)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		art := new(repository.Art)

		if errScan := rows.Scan(&art.Id, &art.Title, &art.Description, &art.Content, &art.AuthorId, &art.Type); errScan != nil {
			return nil, errScan
		}

		result = append(result, art)
	}

	return result, nil
}
