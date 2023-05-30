package column

import "sharkedule/kanban/board"

func (c *Column) Save() error {
	board, err := c.GetParentBoard()
	if err != nil {
		return err
	}
	updateColumnIndexes(board)
	board.Columns[c.Index] = c

	return board.Save()
}

func updateColumnIndexes(board *board.Board) {
	for i, c := range board.Columns {
		c.Index = i
	}
}