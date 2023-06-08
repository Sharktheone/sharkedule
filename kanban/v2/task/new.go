package task

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
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
