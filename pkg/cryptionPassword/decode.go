package cryptionPassword

import (
	"encoding/base64"
	"errors"
	"strings"
)

// error when decoding the password
func Decode(
	dirtyPassword, salt, secondSalt string,
) (string, error) {
	// Убираем пароль с конца пароля
	slice1 := strings.Split(dirtyPassword, secondSalt)
	if len(slice1) != 2 {
		return "", errors.New("error when decoding the password")
	}

	// Декодируем пароль
	pass, err := base64.StdEncoding.DecodeString(slice1[0])
	if err != nil {
		return "", err
	}

	// Убираем соль с начала пароля
	slice2 := strings.Split(string(pass), salt)
	if len(slice2) != 2 {
		return "", errors.New("error when decoding the password")
	}

	return slice2[1], nil
}
