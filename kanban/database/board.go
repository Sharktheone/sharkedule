package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

func CreateBoard(boards []*types2.Board, name string) *types2.Board {
	b := types2.NewBoard(name)
	boards = append(boards, b)
	return b
}

func GetBoard(boards []*types2.Board, uuid string) (*types2.Board, error) {
	for _, b := range boards {
		if b.UUID == uuid {
			return b, nil
		}
	}
	return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
}

func GetBoards(boards []*types2.Board) []*types2.Board {
	return boards
}

func GetBoardNames(boards []*types2.Board) []*namelist.NameList {
	var names []*namelist.NameList
	for _, b := range boards {
		names = append(names, &namelist.NameList{
			Name: b.Name,
			UUID: b.UUID,
		})
	}
	return names
}

func SaveBoard(boards []*types2.Board, b *types2.Board) error {
	for i, b := range boards {
		if b.UUID == b.UUID {
			boards[i] = b
			return nil
		}
	}
	return fmt.Errorf("board with uuid %s does not exist", b.UUID)
}

func SaveBoards(boards []*types2.Board, boardsToSave []*types2.Board) {
	boards = boardsToSave
}

func RemoveColumnFromBoard(board *types2.Board, column string) error {
	for index, c := range board.Columns {
		if c == column {
			board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing column %s not found on board %s", column, board.UUID)
}

func DeleteBoard(boards []*types2.Board, uuid string) error {
	for index, b := range boards {
		if b.UUID == uuid {
			boards = append(boards[:index], boards[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting board %s not found", uuid)
}
