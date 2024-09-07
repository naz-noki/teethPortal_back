package ssoRepository

import "MySotre/pkg/pgDB"

func (s *ssoRepository) SetRefreshToken(userId int, token string) error {
	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO refresh_tokens(
			token,
			user_id
		) VALUES ($1, $2)
	`, token, userId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
