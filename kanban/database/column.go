package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func NewColumn(columns *[]*types.Column, board *types.Board, name string) *types.Column {
	c := types.NewColumn(name)
	board.Columns = append(board.Columns, c.UUID)
	*columns = append(*columns, c)
	return c
}

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

	for i := range board.Columns {
		if i == toIndex {
			board.Columns = append(board.Columns[:i], append([]string{column}, board.Columns[i:]...)...)
			return nil
		}
	} //TODO: Merge into one for loop

	return fmt.Errorf("%v to index %v: index out of range", column, toIndex)
}

func RemoveTaskOnColumn(column *types.Column, task string) error {
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
func RenameColumn(column *types.Column, name string) {
	column.Name = name
}
