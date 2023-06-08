package column

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
	"github.com/google/uuid"
)

func NewColumn(name string) *Column {
	return &Column{
		Column: &types.Column{
			UUID: uuid.New().String(),
			Name: name,
		},
	}
}
