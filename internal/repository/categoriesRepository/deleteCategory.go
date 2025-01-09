package categoriesRepository

import (
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) DeleteCategory(id int) error {
	query := `
		DELETE FROM categories 
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, errExecContext := pgDB.DB.ExecContext(ctx, query, id)
	if errExecContext != nil {
		return errExecContext
	}

	return nil
}
