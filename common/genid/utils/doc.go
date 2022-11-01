package utils

import (
	"time"
	"math/rand"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func docGoRandInt32(maxN int) int {
	return rand.Intn(maxN)
}

