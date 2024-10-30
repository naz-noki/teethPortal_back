package authorsRepository

import "MySotre/pkg/pgDB"

func (ar *authorsRepository) DeleteAuthor(authorId int) error {
	rows, errQuery := pgDB.DB.Query(`
		DELETE FROM authors
		WHERE id = $1;
	`, authorId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
