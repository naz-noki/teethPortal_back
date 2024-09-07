package authorsRepository

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/pgDB"
)

func (a *authorsRepository) UpdatePathToAuthorAvatar(
	avatarPath, authorName string,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE authors 
		SET avatar = $1
		WHERE name = $2
	`, avatarPath, authorName)

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
