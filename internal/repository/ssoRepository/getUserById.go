package ssoRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (s *ssoRepository) GetUserById(id int) (*repository.User, error) {
	user := new(repository.User)

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, login, password, is_admin  
		FROM users 
		WHERE id = $1
	`, id)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&user.Id, &user.Login, &user.Password, &user.IsAdmin); errScan != nil {
			return nil, errScan
		}
	}

	return user, nil
}
