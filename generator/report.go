package generator

import "github.com/Armatorix/SlackUp/model"

func Report() *model.Report {
	return &model.Report{
		Greeting:      Greeting(),
		DoneYesterday: []model.Task{},
		WorkToday:     []model.Task{},
	}
}
