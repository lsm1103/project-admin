package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func docGoRandInt32(maxN int) int {
	return rand.Intn(maxN)
}
