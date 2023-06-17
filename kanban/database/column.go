package kanbandb

import (
	"fmt"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

func NewColumn(columns *[]*types2.Column, board *types2.Board, name string) *types2.Column {
	c := types2.NewColumn(name)
	board.Columns = append(board.Columns, c.UUID)
	*columns = append(*columns, c)
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
	if toIndex < 0 {
		return fmt.Errorf("cannot move column to negative index")
	}

	var l = len(board.Columns) - 1

	for i, c := range board.Columns {
		if c == column {
			board.Columns = append(board.Columns[:i], board.Columns[i+1:]...)
			break
		}
	}

	if toIndex == l {
		board.Columns = append(board.Columns, column)
		return nil
	} else if toIndex > l {
		board.Columns = append(board.Columns, column)
		return fmt.Errorf("%v to index %v: index out of range. Moving to last index", column, toIndex)
	}

	for i, _ := range board.Columns {
		if i == toIndex {
			board.Columns = append(board.Columns[:i], append([]string{column}, board.Columns[i:]...)...)
			return nil
		}
	} //TODO: Merge into one for loop

	return fmt.Errorf("%v to index %v: index out of range", column, toIndex)
}

func RemoveTaskOnColumn(column *types2.Column, task string) error {
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
