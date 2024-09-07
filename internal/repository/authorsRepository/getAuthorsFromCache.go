package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/cacheDB"
	"encoding/json"
)

func (a *authorsRepository) getAuthorsFromCache() []repository.Author {
	// Получаем все данные из кеша
	data, errGetAll := cacheDB.DB.GetAll(repository.TableAuthors)

	if errGetAll != nil {
		return nil
	}
	result := make([]repository.Author, len(data))

	// Демаршалим все данные
	for i := 0; i < len(data); i++ {
		author := new(repository.Author)

		if len(data[i]) < 1 {
			continue
		}

		errUnmarshal := json.Unmarshal(
			[]byte(data[i]),
			author,
		)

		if errUnmarshal != nil {
			continue
		}
		result[i] = *author
	}

	return result
}
