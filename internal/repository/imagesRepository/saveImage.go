package imagesRepository

import (
	"fmt"
	"os"
)

func (u *imagesRepository) SaveImage(
	dirPath, fileName string,
	userId int,
	data *[]byte,
) (string, error) {
	// Проверяем существует ли папка пользователя с изображениями пользователя
	dir := fmt.Sprintf("%s/%d", dirPath, userId)

	if errMkdir := os.Mkdir(dir, 0755); errMkdir != nil && !os.IsExist(errMkdir) {
		return "", errMkdir
	}

	// Создаём файл
	filePath := fmt.Sprintf("%s/%s", dir, fileName)
	file, errOpenFile := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if errOpenFile != nil {
		return "", errOpenFile
	}
	defer file.Close()

	// Пишем данные в файл
	if _, errWrite := file.Write(*data); errWrite != nil {
		return "", errWrite
	}

	return filePath, nil
}
