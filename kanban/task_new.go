package kanban

import "github.com/google/uuid"

func NewTask(name string) *Task {
	return &Task{
		UUID: uuid.New().String(),
		Name: name,
	}
}
