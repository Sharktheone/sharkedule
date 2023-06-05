package types

import "github.com/google/uuid"

func NewBoard(name string) *Board {
	return &Board{
		UUID: uuid.New().String(),
		Name: name,
	}
}
