package kanban

import "github.com/google/uuid"

func NewColumn(name string) *Column {
	return &Column{
		UUID: uuid.New().String(),
		Name: name,
	}
}
