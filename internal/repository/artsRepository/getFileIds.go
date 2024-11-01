package artsRepository

import "MySotre/pkg/pgDB"

func (ar *artsRepository) GetFileIds(artId int) ([]string, error) {
	result := make([]string, 0, 9)

	rows, errQuery := pgDB.DB.Query(`
		SELECT file_id 
		FROM art_files
		WHERE art_id = $1;
	`, artId)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		var fileId string

		if errScan := rows.Scan(&fileId); errScan != nil {
			return nil, errScan
		}

		result = append(result, fileId)
	}

	return result, nil
}
