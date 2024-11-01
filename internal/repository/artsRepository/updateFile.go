package artsRepository

import (
	"MySotre/pkg/minioDB"
	"MySotre/pkg/pgDB"
	"mime/multipart"
)

func (ar *artsRepository) UpdateFile(
	artId int,
	oldFileName string,
	fileHeader *multipart.FileHeader,
) error {
	errRemove := minioDB.Client.Remove(oldFileName, ar.bucketName)

	if errRemove != nil {
		return errRemove
	}

	file, errOpen := fileHeader.Open()
	if errOpen != nil {
		return errOpen
	}
	defer file.Close()

	fileName, errUpload := minioDB.Client.Upload(ar.bucketName, fileHeader.Filename, fileHeader.Size, file)
	if errUpload != nil {
		return errUpload
	}

	rows, errQuery := pgDB.DB.Query(`
		UPDATE art_files 
		SET file_id = $1
		WHERE art_id = $2 AND file_id = $3; 
	`, fileName, artId, oldFileName)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
