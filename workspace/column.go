package workspace

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/column"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

// Column functions
func (w *Workspace) SaveColumn(column *column.Column) error {
	return db.DB.SaveColumn(w.UUID, column.Column)

}

func (w *Workspace) SaveColumns(columns []*column.Column) error {
	var c []*types.Column
	for _, col := range columns {
		c = append(c, col.Column)
	}
	return db.DB.SaveColumns(w.UUID, c)
}

func (w *Workspace) GetColumn(uuid string) (*column.Column, error) {
	return column.Get(w.UUID, uuid)
}

func (w *Workspace) DeleteColumnOnBoard(board, column string) error {
	return db.DB.DeleteColumnOnBoard(w.UUID, board, column)
}

func (w *Workspace) RenameColumn(column, name string) error {
	return db.DB.RenameColumn(w.UUID, column, name)
}

func (w *Workspace) DeleteColumn(uuid string) error {
	return db.DB.DeleteColumn(w.UUID, uuid)
}

func (w *Workspace) MoveColumn(board, uuid string, toIndex int) error {
	return db.DB.MoveColumn(w.UUID, board, uuid, toIndex)
}

func (w *Workspace) NewColumn(board, name string) (*column.Column, error) {
	c, err := db.DB.NewColumn(w.UUID, board, name)
	if err != nil {
		return nil, err
	}
	return &column.Column{Column: c, Workspace: w.UUID}, nil

}
