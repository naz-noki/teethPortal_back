package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (a *authorsRepository) GetAuthors() (*[]repository.Author, error) {
	result := make([]repository.Author, 0, 9)

	// Пробуем получить данные из кэша
	data := a.getAuthorsFromCache()
	if len(data) > 1 {
		return &data, nil
	}

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, name, description, user_id, avatar
		FROM authors
	`)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		author := new(repository.Author)

		if errScan := rows.Scan(&author.Id, &author.Name, &author.Description, &author.UserId, &author.Avatar); errScan != nil {
			return &result, errScan
		}

		result = append(result, *author)
	}

	// Обновляем данные в кэше
	a.setAuthorsInCache(result)

	return &result, nil
}
