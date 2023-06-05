package kanbandb

import types2 "sharkedule/kanban/v2/types"

func CreateBoard(boards []*types2.Board, name string) *types2.Board {
	board := types2.NewBoard(name)
	boards = append(boards, board)
	return board
}
