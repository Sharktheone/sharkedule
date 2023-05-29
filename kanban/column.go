package kanban

import (
	"errors"
	"sharkedule/kanban/column"
)

func (b *Board) GetColumn(uuid string) (*column.Column, error) {
	for _, col := range b.Columns {
		if col.UUID == uuid {
			return col, nil
		}
	}
	return nil, errors.New("column not found")
}

func (b *Board) NewColumn(name string) *column.Column {
	col := column.New(name)
	b.Columns = append(b.Columns, col)
	return col
}
