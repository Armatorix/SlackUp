package model

type Report struct {
	Greeting      Greeting
	DoneYesterday []Task
	WorkToday     []Task
}

type Greeting struct {
	Val string
}

type Task struct {
	Component string
	Action    string
	Status    string
}
