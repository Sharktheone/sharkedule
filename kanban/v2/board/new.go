package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
)

func NewBoard(name string) (*Board, error) {
	b, err := db.DBV2.CreateBoard(name)
	if err != nil {
		return nil, err
	}

	return &Board{Board: b}, nil
}

func (b *Board) NewColumn(name string) (*column.Column, error) {
	c, err := db.DBV2.NewColumn(b.UUID, name)
	if err != nil {
		return nil, err
	}
	return &column.Column{Column: c}, nil
}
