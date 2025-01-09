package globalSearchRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *globalSearchRepository) GetAuthors(searchText string) ([]*repository.Author, error) {
	query := `
		SELECT 
			id, name, description, user_id, avatar_id
		FROM authors 
		WHERE name LIKE $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, errQueryContext := pgDB.DB.QueryContext(ctx, query, "%"+searchText+"%")
	if errQueryContext != nil {
		return nil, errQueryContext
	}
	authors := make([]*repository.Author, 0, 9)

	for rows.Next() {
		author := new(repository.Author)
		errScan := rows.Scan(&author.Id, &author.Name, &author.Description, &author.UserId, &author.AvatarId)

		if errScan != nil {
			return nil, errScan
		}
		authors = append(authors, author)
	}

	if errNext := rows.Err(); errNext != nil {
		return nil, errNext
	}

	return authors, nil
}
