package model

type Report struct {
	Greeting      Greeting
	DoneYesterday []Task
	WorkToday     []Task
}

type Greeting struct {
	Val string
}

type TaskStatus string

const (
	TaskStatusStarted = TaskStatus("started")
	TaskStatusWIP     = TaskStatus("work in progress")
	TaskStatusDone    = TaskStatus("done")
)

type Task struct {
	component string
	status    TaskStatus
	reason    string
}
