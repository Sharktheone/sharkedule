package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/column"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
)

func SaveBoard(boards []*board.Board, b *board.Board) error {
	for i, b := range boards {
		if b.UUID == b.UUID {
			boards[i] = b
			return nil
		}
	}
	return fmt.Errorf("board with uuid %s does not exist", b.UUID)
}

func SaveBoards(boards []*board.Board, boardsToSave []*board.Board) {
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
