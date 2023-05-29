package kanban

import (
	"errors"
)

func (b *Board) GetColumn(uuid string) (*Column, error) {
	for _, col := range b.Columns {
		if col.UUID == uuid {
			return col, nil
		}
	}
	return nil, errors.New("column not found")
}

func (b *Board) NewColumn(name string) *Column {
	col := NewColumn(name)
	b.Columns = append(b.Columns, col)
	return col
}
