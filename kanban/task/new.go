package task

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/google/uuid"
)

func NewTask(name string) *Task {
	return &Task{
		Task: &types.Task{
			UUID: uuid.New().String(),
			Name: name,
		},
	}
}
