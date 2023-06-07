package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
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

func SaveTask(tasks []*task.Task, task *task.Task) error {
	for i, t := range tasks {
		if t.UUID == task.UUID {
			tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task with uuid %s does not exist", task.UUID)
}

func SaveTasks(tasks []*task.Task, tasksToSave []*task.Task) {
	tasks = tasksToSave
}
