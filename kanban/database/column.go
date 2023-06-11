package kanbandb

import (
	"fmt"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

func NewColumn(board *types2.Board, name string) *types2.Column {
	c := types2.NewColumn(name)
	board.Columns = append(board.Columns, c.UUID)
	return c
}

func GetColumn(columns []*types2.Column, uuid string) (*types2.Column, error) {
	for _, c := range columns {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
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

func MoveColumn(board *types2.Board, column string, toIndex int) error {
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

func RemoveTaskFromColumn(column *types2.Column, task string) error {
	for index, t := range column.Tasks {
		if t == task {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while removing task %s not found on column %s", task, column.UUID)
}

func DeleteColumn(columns []*types2.Column, uuid string) error {
	for index, c := range columns {
		if c.UUID == uuid {
			columns = append(columns[:index], columns[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting column %s not found", uuid)
}
func RenameColumn(column *types2.Column, name string) {
	column.Name = name
}
