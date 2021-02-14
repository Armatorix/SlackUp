package generator

import (
	"math/rand"
)

var (
	stumpGreetings = []string{
		"hi",
		"hello",
		"hey everyone",
	}
)

func Greeting() string {
	return stumpGreetings[rand.Intn(len(stumpGreetings))]
}
