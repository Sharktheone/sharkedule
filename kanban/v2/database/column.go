package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
)

func DeleteColumn(columns []*column.Column, uuid string) error {
	for index, c := range columns {
		if c.UUID == uuid {
			columns = append(columns[:index], columns[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting column %s not found", uuid)
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

func GetColumn(columns []*column.Column, uuid string) (*column.Column, error) {
	for _, c := range columns {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
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
