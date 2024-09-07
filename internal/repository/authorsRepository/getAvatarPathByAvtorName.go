package authorsRepository

import "MySotre/pkg/pgDB"

func (a *authorsRepository) GetAvatarPathByAvtorName(name string) (string, error) {
	avatar := ""

	rows, errQuery := pgDB.DB.Query(`
		SELECT avatar FROM authors 
		WHERE name = $1
	`, name)

	if errQuery != nil {
		return avatar, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&avatar); errScan != nil {
			return avatar, errScan
		}
	}

	return avatar, nil
}
