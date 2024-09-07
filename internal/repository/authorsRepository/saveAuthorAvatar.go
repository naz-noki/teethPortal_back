package authorsRepository

import (
	"fmt"
	"os"
	"path"
)

func (a *authorsRepository) SaveAuthorAvatar(
	authorName, dirPath,
	avatarName string, avatarData []byte,
) (string, error) {
	// Получаем путь к файлу и создаём его
	avatar := fmt.Sprintf("%s_-_%s", authorName, avatarName)
	avatarPath := path.Join(dirPath, avatar)
	file, errOpenFile := os.OpenFile(avatarPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if errOpenFile != nil {
		return "", errOpenFile
	}
	defer file.Close()

	// Пишем данные в файл
	_, errWrite := file.Write(avatarData)

	if errWrite != nil {
		return "", errWrite
	}

	return avatarPath, nil
}
