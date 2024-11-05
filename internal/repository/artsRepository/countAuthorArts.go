package artsRepository

import "MySotre/pkg/pgDB"

func (ar *artsRepository) CountAuthorArts(authorId int) (int, error) {
	var counter int

	rows, errQuery := pgDB.DB.Query(`
		SELECT COUNT(id) 
		FROM arts
		WHERE author_id = $1; 
	`, authorId)

	if errQuery != nil {
		return -1, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&counter); errScan != nil {
			return -1, errScan
		}
	}

	return counter, nil
}
