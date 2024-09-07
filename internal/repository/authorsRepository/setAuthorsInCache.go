package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/cacheDB"
	"strconv"
)

func (a *authorsRepository) setAuthorsInCache(authors []repository.Author) {
	for i := 0; i < len(authors); i++ {
		cacheDB.DB.Set(
			repository.TableAuthors,
			strconv.Itoa(authors[i].Id),
			&authors[i],
		)
	}
}
