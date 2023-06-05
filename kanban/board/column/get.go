package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/database/types"
)

func (c *Column) GetParentBoard() (*types.Board, error) {
	return db.DB.GetBoard(c.Board)
}

func GetColumn(board, column string) (*types.Column, error) {
	return db.DB.GetColumn(board, column)
}
