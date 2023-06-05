package types

import "github.com/google/uuid"

func NewBoard(name string) *Board {
	return &Board{
		UUID: uuid.New().String(),
		Name: name,
	}
}

func NewColumn(name string) *Column {
	return &Column{
		UUID: uuid.New().String(),
		Name: name,
	}
}

func NewTask(name string) *Task {
	return &Task{
		UUID: uuid.New().String(),
		Name: name,
	}
}

func NewTag(name string) *Tag {
	return &Tag{
		UUID: uuid.New().String(),
		Name: name,
	}
}
