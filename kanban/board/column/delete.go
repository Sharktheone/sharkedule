package column

import "github.com/Sharktheone/sharkedule/database/db"

func (c *Column) Delete() error {
	board, err := c.GetParentBoard()
	if err != nil {
		return err
	}

	board.Columns = append(board.Columns[:c.Index], board.Columns[c.Index+1:]...)

	return db.DB.SaveBoard(board)
}
