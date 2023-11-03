package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/column"
)

func NewBoard(workspace, name string) (*Board, error) {
	b, err := db.DB.CreateBoard(workspace, name)
	if err != nil {
		return nil, err
	}

	return &Board{Board: b}, nil
}

func (b *Board) NewColumn(name string) (*column.Column, error) {
	c, err := db.DB.NewColumn(b.Workspace, b.UUID, name)
	if err != nil {
		return nil, err
	}
	return &column.Column{Column: c, Workspace: b.Workspace}, nil
}
