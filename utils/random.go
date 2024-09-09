package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandOwner() string {
	return RandString(6)
}

func RandomMoney() int64 {
	return RandInt(0, 1000)
}

func RandomCurrency() string {
	currensies := []string{"EUR", "USD", "UZS"}
	n := len(currensies)
	return currensies[rand.Intn(n)]
}
