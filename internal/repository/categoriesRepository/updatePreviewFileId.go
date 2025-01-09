package categoriesRepository

import (
	"MySotre/pkg/pgDB"
	"context"
	"time"
)

func (t *categoriesRepository) UpdatePreviewFileId(
	id int,
	newPreviewFileId string,
) error {
	query := `
		UPDATE categories 
		SET preview_file_id = $1
		WHERE id = $2 
	`
	params := []interface{}{newPreviewFileId, id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, errExecContext := pgDB.DB.ExecContext(ctx, query, params...)
	if errExecContext != nil {
		return errExecContext
	}

	return nil
}
