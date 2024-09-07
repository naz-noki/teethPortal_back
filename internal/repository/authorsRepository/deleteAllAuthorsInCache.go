package authorsRepository

import (
	"MySotre/internal/repository"
	"MySotre/pkg/cacheDB"
)

func (a *authorsRepository) deleteAllAuthorsInCache() error {
	return cacheDB.DB.Delete(repository.TableAuthors)
}
