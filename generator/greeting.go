package generator

import (
	"math/rand"

	"github.com/Armatorix/SlackUp/model"
)

var (
	stumpGreetings = []string{
		"hi",
		"hello",
		"hey everyone",
	}
)

func Greeting() model.Greeting {
	return model.Greeting{
		Val: stumpGreetings[rand.Intn(len(stumpGreetings))],
	}
}
