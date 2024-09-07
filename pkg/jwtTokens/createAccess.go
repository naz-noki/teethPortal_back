package jwtTokens

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccess(
	secret string,
	userPayload interface{},
	expiresDelay time.Duration,
) (string, error) {
	// Маршалим информацию о владельце токена
	payload, errMarshal := json.Marshal(userPayload)

	if errMarshal != nil {
		return "", nil
	}
	// Создаём токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(payload),
		"exp":  time.Now().Add(expiresDelay).Unix(),
		"iat":  time.Now().Unix(),
	})
	// Подписываем токен
	tokenString, errSignedString := token.SignedString([]byte(secret))

	if errSignedString != nil {
		return "", errSignedString
	}

	return tokenString, nil
}
