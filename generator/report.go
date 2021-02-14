package generator

import "github.com/Armatorix/SlackUp/model"

func Report() *model.Report {
	yesterday, today := Tasks()
	return &model.Report{
		Greeting:      Greeting(),
		DoneYesterday: yesterday,
		WorkToday:     today,
	}
}
