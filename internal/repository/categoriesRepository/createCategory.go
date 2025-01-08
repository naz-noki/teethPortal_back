package categoriesRepository

import (
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) CreateCategory(
	name, description,
	previewFileId string,
) error {
	query := `INSERT INTO categories (name, description, preview_file_id) VALUES (
		$1, $2, $3
	)`
	params := []interface{}{name, description, previewFileId}

	ctx, cncl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cncl()

	_, errExecContext := pgDB.DB.ExecContext(ctx, query, params...)
	if errExecContext != nil {
		return errExecContext
	}
	return nil
}
