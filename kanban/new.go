package kanban

import "github.com/google/uuid"

func New(name string) *Board {
	return &Board{
		UUID: uuid.New().String(),
		Name: name,
	}
}
