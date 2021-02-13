package generator

import (
	"math/rand"

	"github.com/Armatorix/SlackUp/model"
)

var (
	greetingsList = []string{
		"hi",
		"hello",
		"hey everyone",
	}
)

func Greeting() model.Greeting {
	return model.Greeting{
		Val: greetingsList[rand.Intn(len(greetingsList))],
	}
}
