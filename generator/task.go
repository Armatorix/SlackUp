package generator

import (
	"math/rand"

	"github.com/Armatorix/SlackUp/model"
	"github.com/Armatorix/SlackUp/x/xmath/xrand"
)

// $status $action for $component XD
var stumpComponents = []string{
	"new feature",
	"task from yesterday",
	"new bug",
	"network bug",
}

var stumpActions = []string{
	"doing tests",
	"writing automation tests",
	"testing manually",
	"implementation",
	"merged ticket",
}

var stumpStatus = []string{
	"started",
	"continuing",
	"wip",
	"finished",
	"done",
}

func TaskN(task, action, status int) model.Task {
	return model.Task{
		Component: stumpComponents[task%len(stumpComponents)],
		Action:    stumpActions[action%len(stumpActions)],
		Status:    stumpStatus[status%len(stumpActions)],
	}
}

func TasksN(n []int) []model.Task {
	tasks := make([]model.Task, len(n))
	for i, v := range n {
		tasks[i] = TaskN(v, rand.Int(), rand.Int())
	}
	return tasks
}

func Tasks() (today, yesterday []model.Task) {
	subPerm1, subPerm2 := subsetUniquePermutation()
	return TasksN(subPerm1), TasksN(subPerm2)
}

func subsetUniquePermutation() ([]int, []int) {
	totalElements := xrand.Intnn(2, len(stumpComponents))
	r := rand.Perm(len(stumpComponents))[:totalElements]
	split := xrand.Intnn(1, totalElements-1)
	return r[split:], r[:split]
}
