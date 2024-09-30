package domain

import (
	"math/rand"
)

func Rand(max int) int {
	min := 0

	return rand.Intn(max-min) + min
}
