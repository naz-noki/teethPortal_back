package artsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (ar *artsRepository) GetAuthorArts(authorId int) ([]*repository.Art, error) {
	result := make([]*repository.Art, 0, 9)

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, title, description, content, author_id, type 
		FROM arts
		WHERE author_id = $1; 
	`, authorId)

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
