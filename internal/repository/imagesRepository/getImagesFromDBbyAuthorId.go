package imagesRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (i *imagesRepository) GetImagesFromDBbyAuthorId(authorId int) (*[]repository.Image, error) {
	images := make([]repository.Image, 0, 9)

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, path, title, description, 
		user_id, author_id, created_at 
		FROM images
		WHERE author_id = $1
	`, authorId)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		image := new(repository.Image)

		if errScan := rows.Scan(
			&image.Id, &image.Path, &image.Title, &image.Description,
			&image.UserId, &image.AuthorId, &image.CreatedAt,
		); errScan != nil {
			return &images, errScan
		}

		images = append(images, *image)
	}

	return &images, nil
}
