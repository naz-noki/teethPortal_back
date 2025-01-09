package categoriesRepository

import (
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) UpdateCategory(
	id int,
	name, description string,
) error {
	query := `
		UPDATE categories 
		SET name = $1, description = $2
		WHERE id = $3 
	`
	params := []interface{}{name, description, id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, errExecContext := pgDB.DB.ExecContext(ctx, query, params...)
	if errExecContext != nil {
		return errExecContext
	}

	return nil
}
