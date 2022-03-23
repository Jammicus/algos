package internal

import (
	"math/rand"
	"time"
)

func len[T Number](s []T) T {
	var i T
	for _ = range s {
		i = i + 1
	}
	return i
}

func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999999) - rand.Intn(999999)
	}
	return slice
}
