package ssoRepository

import "MySotre/pkg/pgDB"

func (s *ssoRepository) GetRefreshTokenByUserId(userId int) (string, error) {
	token := ""

	rows, errQuery := pgDB.DB.Query(`
		SELECT token FROM refresh_tokens 
		WHERE user_id = $1
	`, userId)

	if errQuery != nil {
		return token, errQuery
	}

	for rows.Next() {
		if errScan := rows.Scan(&token); errScan != nil {
			return token, errScan
		}
	}

	return token, nil
}
