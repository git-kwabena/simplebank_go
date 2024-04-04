package utils

import (
	"math/rand"
	"time"
)

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	//seed is deprecated
	//rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	owner := make([]byte, n)
	for i := range owner {
		owner[i] = alphabets[rng.Intn(len(alphabets))]
	}
	return string(owner)
}
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomOwner() string {
	return RandomString(6)
}
func RandomCurrency(s string) string {
	currencies := []string{"USD", "EUR", "JPY", "GBP", "AUD", "CAD"}
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return currencies[rng.Intn(len(currencies))]
}

func RandomBalance() int64 {

	return RandomInt(0, 1000)
}
