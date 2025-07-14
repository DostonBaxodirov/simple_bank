package utils

import (
	"math/rand/v2"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int64N(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.IntN(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandOwner() string {
	return RandomString(8)
}

func RandBalance() int64 {
	return RandomInt(0, 10000)
}

func RandCurrency() string {
	currencies := []string{"EUR", "USD", "UZS"}
	n := len(currencies)
	return currencies[rand.IntN(n)]
}
