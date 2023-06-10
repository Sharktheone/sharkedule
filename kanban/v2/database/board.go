package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func CreateBoard(boards []*types.Board, name string) *types.Board {
	b := types.NewBoard(name)
	boards = append(boards, b)
	return b
}

func GetBoard(boards []*types.Board, uuid string) (*types.Board, error) {
	for _, b := range boards {
		if b.UUID == uuid {
			return b, nil
		}
	}
	return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
}

func GetBoards(boards []*types.Board) []*types.Board {
	return boards
}

func GetBoardNames(boards []*types.Board) []*namelist.NameList {
	var names []*namelist.NameList
	for _, b := range boards {
		names = append(names, &namelist.NameList{
			Name: b.Name,
			UUID: b.UUID,
		})
	}
	return names
}

func SaveBoard(boards []*types.Board, b *types.Board) error {
	for i, b := range boards {
		if b.UUID == b.UUID {
			boards[i] = b
			return nil
		}
	}
	return fmt.Errorf("board with uuid %s does not exist", b.UUID)
}

func SaveBoards(boards []*types.Board, boardsToSave []*types.Board) {
	boards = boardsToSave
}

func RemoveColumnFromBoard(board *types.Board, column string) error {
	for index, c := range board.Columns {
		if c == column {
			board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing column %s not found on board %s", column, board.UUID)
}

func DeleteBoard(boards []*types.Board, uuid string) error {
	for index, b := range boards {
		if b.UUID == uuid {
			boards = append(boards[:index], boards[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting board %s not found", uuid)
}
