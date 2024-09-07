package imagesRepository

import "MySotre/pkg/pgDB"

func (i *imagesRepository) UpdateImageData(
	imageId int,
	newTitle, newDescription string,
) error {
	rows, errQuery := pgDB.DB.Query(`
		UPDATE images
		SET title = $1, description = $2
		WHERE id = $3
	`, newTitle, newDescription, imageId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
