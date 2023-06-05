package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/database/types"
)

func (c *Column) Save() error {
	board, err := c.GetParentBoard()
	if err != nil {
		return err
	}
	updateColumnIndexes(board)
	col, err := c.Convert()
	if err != nil {
		return err
	}
	board.Columns[c.Index] = col

	return db.DB.SaveBoard(board)
}

func updateColumnIndexes(board *types.Board) {
	for i, c := range board.Columns {
		c.Index = i
	}
}
