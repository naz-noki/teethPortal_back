package categoriesRepository

import (
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) AddAuthor(authorId, categoryId int) error {
	query := `INSERT INTO author_category (author_id, category_id) VALUES (
		$1, $2
	)`
	params := []interface{}{authorId, categoryId}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, errExecContext := pgDB.DB.ExecContext(ctx, query, params...)
	if errExecContext != nil {
		return errExecContext
	}
	return nil
}
