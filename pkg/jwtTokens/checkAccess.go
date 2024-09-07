package jwtTokens

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CheckAccess(
	token, secret string,
	userPayload interface{},
) error {
	// Парсим токен
	t, errParse := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})

	if errParse != nil {
		return errParse
	}

	// Проверяем валидность токена
	if !t.Valid {
		return errors.New("error checking the validity of the token")
	}

	// Получаем claims из токена
	claims, okClaims := t.Claims.(jwt.MapClaims)

	if !okClaims {
		return errors.New("an error occurred while getting claims from the token")
	}

	// Получаем информацию о владельце токена из claims
	userData, ok := claims["user"].(string)

	if !ok {
		return errors.New("an error occurred while getting claims from the token")
	}

	// Парсим информацию о владельце токена
	return json.Unmarshal([]byte(userData), userPayload)
}
