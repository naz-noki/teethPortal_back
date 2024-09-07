package cryptionPassword

import "encoding/base64"

func Encode(
	password, salt, secondSalt string,
) string {
	// Добавляем соль в начало пароля
	password = salt + password
	// Кодируем пароль
	password = base64.StdEncoding.EncodeToString([]byte(password))
	// Добавляем соль в конец пароля
	password = password + secondSalt

	return password
}
