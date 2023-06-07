package kanbandb

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
)

func CreateBoard(boards []*board.Board, name string) *board.Board {
	b := board.NewBoard(name)
	boards = append(boards, b)
	return b
}
