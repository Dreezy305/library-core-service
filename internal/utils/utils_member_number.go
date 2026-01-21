package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateMemberNumber() string {
	year := time.Now().Year()
	suffix := randString(5)
	return fmt.Sprintf("MBR-%d-%s", year, suffix)
}
