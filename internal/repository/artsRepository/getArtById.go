package artsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (artsRepository *artsRepository) GetArtById(id int) (*repository.Art, error) {
	art := new(repository.Art)

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, title, description, content, author_id, type 
		FROM arts
		WHERE id = $1;  
	`, id)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&art.Id, &art.Title, &art.Description, &art.Content, &art.AuthorId, &art.Type); errScan != nil {
			return nil, errScan
		}
	}

	return art, nil
}
