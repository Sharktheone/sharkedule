package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/database/types"
)

func (c *Column) Move(toIndex int) error {
	board, err := c.GetParentBoard()
	if err != nil {
		return err
	}

	board.Columns = append(board.Columns[:c.Index], board.Columns[c.Index+1:]...)

	col, err := c.Convert()
	if err != nil {
		return err
	}

	board.Columns = append(board.Columns[:toIndex], append([]*types.Column{col}, board.Columns[toIndex:]...)...)

	return db.DB.SaveBoard(board)
}
