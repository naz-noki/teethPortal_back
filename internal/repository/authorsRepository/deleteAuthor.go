package authorsRepository

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/pgDB"
)

func (a *authorsRepository) DeleteAuthor(id int) error {
	rows, errQuery := pgDB.DB.Query(`
		DELETE FROM authors
		WHERE id = $1
	`, id)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()
	// Удаляем данные из кеша, т.к. они теперь не актуальны
	if errDeleteAllAuthorsInCache := a.deleteAllAuthorsInCache(); errDeleteAllAuthorsInCache != nil {
		logger.Log.Debug(errDeleteAllAuthorsInCache.Error())
	}

	return nil
}
