package model

import (
	"fmt"
	"strings"

	"github.com/Armatorix/SlackUp/x"
)

type Report struct {
	Greeting      string
	DoneYesterday []Task
	WorkToday     []Task
}

func (r Report) String() string {
	return fmt.Sprintf("%s! Yesterday I:\n%s\n and today I will:\n%s\n",
		strings.Title(r.Greeting),
		strings.Join(x.StringerToString(r.DoneYesterday), "\n"),
		strings.Join(x.StringerToString(r.WorkToday), "\n"))
}

type Greeting string

type Task struct {
	Component string
	Action    string
	Status    string
}

func (t Task) String() string {
	return fmt.Sprintf("* %s %s for %s", t.Status, t.Action, t.Component)
}
