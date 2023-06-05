package task

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/database/types"
)

func (t *Task) GetParentBoard() (*types.Board, error) {
	return db.DB.GetBoard(t.Board)
}

func (t *Task) GetParentColumn() (*types.Column, error) {
	return db.DB.GetColumn(t.Board, t.Column)
}
