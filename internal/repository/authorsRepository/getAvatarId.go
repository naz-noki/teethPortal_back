package authorsRepository

import "MySotre/pkg/pgDB"

func (ar *authorsRepository) GetAvatarId(authorId int) (string, error) {
	var avatarId string

	rows, errQuery := pgDB.DB.Query(`
		SELECT avatar_id 
		FROM authors
		WHERE id = $1;
	`, authorId)

	if errQuery != nil {
		return "", errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&avatarId); errScan != nil {
			return "", errScan
		}
	}

	return avatarId, nil
}
