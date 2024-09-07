package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/pgDB"
)

func (a *authorsRepository) GetAuthorById(id int) (*repository.Author, error) {
	author := new(repository.Author)

	// Пробуем получить данные из кэша
	data := a.getAuthorByIdFromCache(id)
	if data != nil {
		return data, nil
	}

	rows, errQuery := pgDB.DB.Query(`
		SELECT id, name, description, user_id, avatar 
		FROM authors 
		WHERE id = $1
	`, id)

	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		if errScan := rows.Scan(&author.Id, &author.Name, &author.Description, &author.UserId, &author.Avatar); errScan != nil {
			return nil, errScan
		}
	}

	// Обновляем данные в кэше
	a.setAuthorsInCache([]repository.Author{*author})

	return author, nil
}
