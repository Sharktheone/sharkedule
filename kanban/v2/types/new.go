package types

import (
	"github.com/google/uuid"
)

func NewColumn(name string) *Column {
	return &Column{
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
