package authorsRepository

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/pgDB"
	"log"
)

func (a *authorsRepository) UpdateAuthor(
	name, description string,
	id int,
) error {
	log.Println(name, description, id)

	rows, errQuery := pgDB.DB.Query(`
		UPDATE authors 
		SET name = $1, description = $2
		WHERE id = $3 
	`, name, description, id)

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
