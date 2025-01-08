package categoriesRepository

import (
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) DeleteAuthor(authorId, categoryId int) error {
	query := `
		DELETE FROM author_category 
		WHERE
			author_id = $1 AND category_id = $2
	`
	params := []interface{}{authorId, categoryId}

	ctx, cncl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cncl()

	_, errExecContext := pgDB.DB.ExecContext(ctx, query, params...)
	if errExecContext != nil {
		return errExecContext
	}
	return nil
}
