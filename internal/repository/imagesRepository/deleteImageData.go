package imagesRepository

import "MySotre/pkg/pgDB"

func (i *imagesRepository) DeleteImageData(imageId int, imagePath string) error {
	rows, errQuery := pgDB.DB.Query(`
		DELETE FROM images 
		WHERE path = $1 AND id = $2
	`, imagePath, imageId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
