package models

import (
	"encoding/json"

	"github.com/golkhandani/taskmate/exceptions"
)

type Task struct {
	ID     int64
	Title  string
	IsDone bool
}

func LoadTasksFromBytes(content []byte) []Task {
	tasks := make([]Task, 0)
	if len(content) == 0 {
		return tasks
	}
	err := json.Unmarshal(content, &tasks)
	exceptions.HandleErr(err)
	return tasks
}
