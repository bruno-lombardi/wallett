package generators

import (
	"fmt"
	"math/rand"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ID(prefix string) string {
	id := fmt.Sprintf("%v_%v", prefix, RandomString(10))
	return id
}

func FloatsInRange(min, max float64, n int) []float64 {
	values := make([]float64, n)
	for i := range values {
		values[i] = min + rand.Float64()*(max-min)
	}
	return values
}
