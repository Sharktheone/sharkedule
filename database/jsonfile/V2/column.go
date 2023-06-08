package jsonfileV2

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
	kanbandb "github.com/Sharktheone/sharkedule/kanban/v2/database"
)

func (J *JSONFile) GetColumn(uuid string) (*column.Column, error) {
	return kanbandb.GetColumn(J.db.Columns, uuid)
}

func (J *JSONFile) SaveColumn(column *column.Column) error {
	if err := kanbandb.SaveColumn(J.db.Columns, column); err != nil {
		return fmt.Errorf("failed saving column: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveColumns(columns []*column.Column) error {
	kanbandb.SaveColumns(J.db.Columns, columns)
	return J.Save()
}

func (J *JSONFile) DeleteColumn(uuid string) error {
	if err := kanbandb.DeleteColumn(J.db.Columns, uuid); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) RemoveColumnFromBoard(boardUUID string, columnUUID string) error {
	b, err := kanbandb.GetBoard(J.db.Boards, boardUUID)
	if err != nil {
		return err
	}
	if err := kanbandb.RemoveColumnFromBoard(b, columnUUID); err != nil {
		return err
	}
	return J.Save()
}
