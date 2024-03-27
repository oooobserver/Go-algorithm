package util

import (
	"math/rand"
	"strings"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

const letters = "abcdefghijklmnopqrstuvwxyz"

// Inclusive
func RandomInt(min, max int) int {
	return min + rng.Intn(max-min+1)
}

// Return a int, can be negative
func JustRandomInt() int {
	res := rng.Int()
	sign := rng.Int()
	if res < sign {
		return -1 * res
	}

	return res
}

func RandomString(n int) string {
	var sb strings.Builder
	l := len(letters)

	for i := 0; i < n; i++ {
		char := letters[rng.Intn(l)]
		sb.WriteByte(char)
	}
	return sb.String()
}
