package categoriesRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) GetCategoryById(id int) (*repository.Category, error) {
	query := `
		SELECT 
			categories.id, categories.name, 
			categories.description, author_category.author_id
		FROM categories 
		INNER JOIN author_category ON categories.id = author_category.category_id
		WHERE categories.id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rows, errQueryContext := pgDB.DB.QueryContext(ctx, query, id)
	if errQueryContext != nil {
		return nil, errQueryContext
	}
	defer rows.Close()

	authorsIds := make([]int, 0, 9)
	category := new(repository.Category)

	for rows.Next() {
		var authorId int
		errScan := rows.Scan(&category.Id, &category.Name, &category.Description, &authorId)

		if errScan != nil {
			return nil, errScan
		}
		authorsIds = append(authorsIds, authorId)
	}

	if errNext := rows.Err(); errNext != nil {
		return nil, errNext
	}
	category.AuthorsIds = authorsIds

	return category, nil
}
