package artsRepository

import (
	"MySotre/pkg/minioDB"
	"MySotre/pkg/pgDB"
	"mime/multipart"
)

func (ar *artsRepository) SaveFile(
	artId int,
	fileHeader *multipart.FileHeader,
) error {
	file, errOpen := fileHeader.Open()
	if errOpen != nil {
		return errOpen
	}
	defer file.Close()

	fileId, errUpload := minioDB.Client.Upload(ar.bucketName, fileHeader.Filename, fileHeader.Size, file)
	if errUpload != nil {
		return errUpload
	}

	rows, errQuery := pgDB.DB.Query(`
		INSERT INTO art_files (
			art_id, file_id
		) VALUES (
			$1, $2
		);
	`, artId, fileId)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
