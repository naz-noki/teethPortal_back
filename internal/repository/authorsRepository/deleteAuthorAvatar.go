package authorsRepository

import "os"

func (a *authorsRepository) DeleteAuthorAvatar(path string) error {
	errRemove := os.Remove(path)

	if errRemove != nil {
		return errRemove
	}

	return nil
}
