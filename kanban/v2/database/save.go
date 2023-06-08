package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

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

func SaveColumn(columns []*column.Column, column *column.Column) error {
	for i, c := range columns {
		if c.UUID == column.UUID {
			columns[i] = column
			return nil
		}
	}
	return fmt.Errorf("column with uuid %s does not exist", column.UUID)
}

func SaveColumns(columns []*column.Column, columnsToSave []*column.Column) {
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
