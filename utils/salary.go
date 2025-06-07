package utils

import (
	"math/rand"
	"time"
)

func RandomSalary() int {
	rand.Seed(time.Now().UnixNano())
	min := 1_000_000
	max := 20_000_000
	return rand.Intn((max-min+1)/1000)*1000 + min
}
