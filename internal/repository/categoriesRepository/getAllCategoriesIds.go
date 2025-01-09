package categoriesRepository

import (
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) GetAllCategoriesIds() ([]int, error) {
	query := ` 
		SELECT id FROM categories
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, errQueryContext := pgDB.DB.QueryContext(ctx, query)
	if errQueryContext != nil {
		return nil, errQueryContext
	}
	defer rows.Close()

	categoriesIds := make([]int, 0, 9)

	for rows.Next() {
		var id int

		if errScan := rows.Scan(&id); errScan != nil {
			return nil, errScan
		}

		categoriesIds = append(categoriesIds, id)
	}

	if errNext := rows.Err(); errNext != nil {
		return nil, errNext
	}

	return categoriesIds, nil
}
