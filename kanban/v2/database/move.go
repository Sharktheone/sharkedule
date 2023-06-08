package kanbandb

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func MoveColumn(board *types.Board, column string, toIndex int) error {
	var (
		deleted  = false
		inserted = false
	)
	for index, c := range board.Columns {
		if c == column {
			board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)
			if inserted {
				return nil
			}
			deleted = true
		}
		if index == toIndex {
			board.Columns = append(board.Columns, c)
			if deleted {
				return nil
			}
			inserted = true
		}
	}
	return nil

}
