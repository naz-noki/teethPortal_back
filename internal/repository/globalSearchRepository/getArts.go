package globalSearchRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *globalSearchRepository) GetArts(searchText string) ([]*repository.Art, error) {
	query := `
		SELECT 
			id, title, description, content, author_id, type
		FROM arts 
		WHERE title LIKE $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, errQueryContext := pgDB.DB.QueryContext(ctx, query, "%"+searchText+"%")
	if errQueryContext != nil {
		return nil, errQueryContext
	}
	arts := make([]*repository.Art, 0, 9)

	for rows.Next() {
		art := new(repository.Art)
		errScan := rows.Scan(&art.Id, &art.Title, &art.Description, &art.Content, &art.AuthorId, &art.Type)

		if errScan != nil {
			return nil, errScan
		}
		arts = append(arts, art)
	}

	if errNext := rows.Err(); errNext != nil {
		return nil, errNext
	}

	return arts, nil
}
