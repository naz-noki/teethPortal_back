package jwtTokens

import (
	"math/rand"
	"strings"
	"time"
)

func CreateRefresh(size int) string {
	result := make([]string, size)
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")

	rand.Seed(time.Now().UnixMilli())

	for i := 0; i < size; i++ {
		idx := rand.Intn(26)
		el := alphabet[idx]

		if i%3 == 0 {
			el = strings.ToUpper(el)
		}

		result[i] = el
	}

	return strings.Join(result, "")
}
