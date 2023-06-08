package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func DeleteBoard(boards []*board.Board, uuid string) error {
	for index, b := range boards {
		if b.UUID == uuid {
			boards = append(boards[:index], boards[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting board %s not found", uuid)
}

func DeleteColumn(columns []*column.Column, uuid string) error {
	for index, c := range columns {
		if c.UUID == uuid {
			columns = append(columns[:index], columns[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting column %s not found", uuid)
}

func DeleteTask(tasks []*task.Task, uuid string) error {
	for index, t := range tasks {
		if t.UUID == uuid {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting task %s not found", uuid)
}

func DeleteTag(tags []*types.Tag, tag string) error {
	for index, t := range tags {
		if t.UUID == tag {
			tags = append(tags[:index], tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting tag %s not found", tag)
}

func RemoveColumnFromBoard(board *board.Board, column string) error {
	for index, c := range board.Columns {
		if c == column {
			board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing column %s not found on board %s", column, board.UUID)
}

func RemoveTaskFromColumn(column *column.Column, task string) error {
	for index, t := range column.Tasks {
		if t == task {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing task %s not found on column %s", task, column.UUID)
}

func RemoveTagFromTask(task *task.Task, tag string) error {
	for index, t := range task.Tags {
		if t == tag {
			task.Tags = append(task.Tags[:index], task.Tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing tag %s not found on task %s", tag, task.UUID)
}
