package types

import (
	"github.com/google/uuid"
)

func NewTag(name string) *Tag {
	return &Tag{
		UUID: uuid.New().String(),
		Name: name,
	}
}

func NewBoard(name string) *Board {
	return &Board{
		UUID: uuid.New().String(),
		Name: name,
	}
}
