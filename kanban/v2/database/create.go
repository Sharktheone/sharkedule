package kanbandb

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func CreateBoard(boards []*types.Board, name string) *types.Board {
	b := types.NewBoard(name)
	boards = append(boards, b)
	return b
}
