package board

import (
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

type Board struct {
	*types.Board
	Workspace string
}

func (b *Board) GetUUID() string {
	return b.UUID
}

func (b *Board) GetName() string {
	return b.Name
}

func (b *Board) GetType() elements.Type {
	return elements.Board
}

func (b *Board) GetFields() []field.Field {
	return nil //TODO
}
