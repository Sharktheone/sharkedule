package board

import (
	"errors"
	"sharkedule/kanban/board/column"
)

func (b *Board) GetColumn(uuid string) (*column.Column, error) {
	for _, col := range b.Columns {
		if col.UUID == uuid {
			return col, nil
		}
	}
	return nil, errors.New("column not found")
}

func (b *Board) NewColumn(name string) (*column.Column, error) {
	col := column.NewColumn(name)
	b.Columns = append(b.Columns, col)
	if err := b.Save(); err != nil {
		return nil, err
	}
	return col, nil
}
