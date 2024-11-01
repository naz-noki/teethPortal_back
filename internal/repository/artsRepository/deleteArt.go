package artsRepository

import "MySotre/pkg/pgDB"

func (ar *artsRepository) DeleteArt(artId int) error {
	rows, errQuery := pgDB.DB.Query(`
		DELETE FROM arts 
		WHERE id = $1; 
	`, artId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
