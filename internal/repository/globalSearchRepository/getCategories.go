package globalSearchRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *globalSearchRepository) GetCategories(searchText string) ([]*repository.Category, error) {
	query := `
		SELECT 
			id, name, description, preview_file_id
		FROM categories 
		WHERE name LIKE $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, errQueryContext := pgDB.DB.QueryContext(ctx, query, "%"+searchText+"%")
	if errQueryContext != nil {
		return nil, errQueryContext
	}
	categories := make([]*repository.Category, 0, 9)

	for rows.Next() {
		category := new(repository.Category)
		errScan := rows.Scan(&category.Id, &category.Name, &category.Description, &category.PreviewFileId)

		if errScan != nil {
			return nil, errScan
		}
		categories = append(categories, category)
	}

	if errNext := rows.Err(); errNext != nil {
		return nil, errNext
	}

	return categories, nil
}
