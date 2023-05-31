package column

import (
	"sharkedule/database/db"
	"sharkedule/database/types"
)

func (c *Column) GetParentBoard() (*types.Board, error) {
	return db.DB.GetBoard(c.Board)
}

func GetColumn(board, column string) (*types.Column, error) {
	return db.DB.GetColumn(board, column)
}
