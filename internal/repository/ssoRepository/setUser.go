package ssoRepository

import "MySotre/pkg/pgDB"

func (s *ssoRepository) SetUser(
	login, password string,
	isAdmin bool,
) error {
	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO users(
			login,
			password,
			is_admin
		) VALUES ($1, $2, $3)
	`, login, password, isAdmin)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
