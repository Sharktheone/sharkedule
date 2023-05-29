package column

import "sharkedule/kanban"

func (c *Column) Save() error {
	board, err := c.GetParentBoard()
	if err != nil {
		return err
	}
	updateIndexes(board)
	board.Columns[c.Index] = c

	return board.Save()
}

func updateIndexes(board *kanban.Board) {
	for i, c := range board.Columns {
		c.Index = i
	}
}
