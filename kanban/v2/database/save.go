package kanbandb

import (
	"fmt"
	"sharkedule/kanban/KTypes/namelist"
	types2 "sharkedule/kanban/v2/types"
)

func SaveBoard(boards []*types2.Board, board *types2.Board) error {
	for i, b := range boards {
		if b.UUID == board.UUID {
			boards[i] = board
			return nil
		}
	}
	return fmt.Errorf("board with uuid %s does not exist", board.UUID)
}

func SaveBoards(boards []*types2.Board, boardsToSave []*types2.Board) {
	boards = boardsToSave
}

func SaveColumn(columns []*types2.Column, column *types2.Column) error {
	for i, c := range columns {
		if c.UUID == column.UUID {
			columns[i] = column
			return nil
		}
	}
	return fmt.Errorf("column with uuid %s does not exist", column.UUID)
}

func SaveColumns(columns []*types2.Column, columnsToSave []*types2.Column) {
	columns = columnsToSave
}

func SaveTask(tasks []*types2.Task, task *types2.Task) error {
	for i, t := range tasks {
		if t.UUID == task.UUID {
			tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task with uuid %s does not exist", task.UUID)
}

func SaveTasks(tasks []*types2.Task, tasksToSave []*types2.Task) {
	tasks = tasksToSave
}

func CreateBoard(boards []*types2.Board, name string) *types2.Board {
	board := types2.NewBoard(name) // TODO: add func
	boards = append(boards, board)
	return board
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

func GetColumn(columns []*types2.Column, uuid string) (*types2.Column, error) {
	for _, c := range columns {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
}

func GetTask(tasks []*types2.Task, uuid string) (*types2.Task, error) {
	for _, t := range tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with uuid %s does not exist", uuid)
}
