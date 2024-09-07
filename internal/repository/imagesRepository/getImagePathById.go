package imagesRepository

import "MySotre/pkg/pgDB"

func (i *imagesRepository) GetImagePathById(id int) (string, error) {
	imagePath := ""

	rows, errQuery := pgDB.DB.Query(`
		SELECT path FROM images
		WHERE id = $1
	`, id)

	if errQuery != nil {
		return "", errQuery
	}

	for rows.Next() {
		if errScan := rows.Scan(&imagePath); errScan != nil {
			return "", errScan
		}
	}

	return imagePath, nil
}
