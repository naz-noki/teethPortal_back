package authorsRepository

import "MySotre/pkg/pgDB"

func (ar *authorsRepository) UpdateAvatarId(
	authorId int,
	avatarId string,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE authors SET avatar_id = $1
		WHERE id = $2;
	`, avatarId, authorId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
