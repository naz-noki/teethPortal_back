package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/cacheDB"
	"encoding/json"
	"strconv"
)

func (a *authorsRepository) getAuthorByIdFromCache(id int) *repository.Author {
	data, errGet := cacheDB.DB.Get(
		repository.TableAuthors,
		strconv.Itoa(id),
	)

	if errGet != nil {
		return nil
	}

	author := new(repository.Author)
	errUnmarshal := json.Unmarshal([]byte(data), author)

	if errUnmarshal != nil {
		return nil
	}

	return author
}
