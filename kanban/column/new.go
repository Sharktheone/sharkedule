package column

import "github.com/google/uuid"

func New(name string) *Column {
	return &Column{
		UUID: uuid.New().String(),
		Name: name,
	}
}
