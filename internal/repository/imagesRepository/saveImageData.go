package imagesRepository

import "MySotre/pkg/pgDB"

func (I *imagesRepository) SaveImageData(
	userId, authorId int,
	path, title,
	description, createdAt string,
) error {
	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO images (
			path,
			title,
			description,
			user_id,
			author_id,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`, path, title, description, userId, authorId, createdAt)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
