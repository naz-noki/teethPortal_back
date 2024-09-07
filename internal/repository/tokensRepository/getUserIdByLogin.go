package tokensRepository

import "MySotre/pkg/pgDB"

func (t *tokensRepository) GetUserIdByLogin(login string) (int, error) {
	id := -1

	rows, errQuery := pgDB.DB.Query(`
		SELECT id FROM users 
		WHERE login = $1
	`, login)

	if errQuery != nil {
		return id, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&id); errScan != nil {
			return id, errScan
		}
	}

	return id, nil
}
