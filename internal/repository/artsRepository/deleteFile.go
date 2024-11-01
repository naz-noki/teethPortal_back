package artsRepository

import (
	"MySotre/pkg/minioDB"
	"MySotre/pkg/pgDB"
)

func (ar *artsRepository) DeleteFile(fileName string) error {
	errRemove := minioDB.Client.Remove(fileName, ar.bucketName)

	if errRemove != nil {
		return errRemove
	}

	rows, errQuery := pgDB.DB.Query(`
		DELETE FROM art_files 
		WHERE file_id = $1; 
	`, fileName)

	if errQuery != nil {
		return errQuery
	}
	defer rows.Close()

	return nil
}
