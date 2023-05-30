package column

import (
	"sharkedule/kanban/board"
)

func (c *Column) GetParentBoard() (*board.Board, error) {
	return board.GetBoard(c.Board)
}
