package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func GetColumn(columns []*types.Column, uuid string) (*types.Column, error) {
	for _, c := range columns {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
}

func SaveColumn(columns []*types.Column, column *types.Column) error {
	for i, c := range columns {
		if c.UUID == column.UUID {
			columns[i] = column
			return nil
		}
	}
	return fmt.Errorf("column with uuid %s does not exist", column.UUID)
}

func SaveColumns(columns []*types.Column, columnsToSave []*types.Column) {
	columns = columnsToSave
}

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

func RemoveTaskFromColumn(column *types.Column, task string) error {
	for index, t := range column.Tasks {
		if t == task {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing task %s not found on column %s", task, column.UUID)
}

func DeleteColumn(columns []*types.Column, uuid string) error {
	for index, c := range columns {
		if c.UUID == uuid {
			columns = append(columns[:index], columns[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting column %s not found", uuid)
}
