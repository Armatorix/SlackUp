package xrand

import "math/rand"

func Intnn(min, max int) int {
	if min > max {
		panic("invalid argument to Intnn")
	}
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}
