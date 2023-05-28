package task

import "github.com/google/uuid"

func New(name string) *Task {
	return &Task{
		UUID: uuid.New().String(),
		Name: name,
	}
}
