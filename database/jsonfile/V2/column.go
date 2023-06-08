package jsonfileV2

import (
	"fmt"
	kanbandb "github.com/Sharktheone/sharkedule/kanban/v2/database"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func (J *JSONFile) GetColumn(uuid string) (*types.Column, error) {
	return kanbandb.GetColumn(J.db.Columns, uuid)
}

func (J *JSONFile) SaveColumn(column *types.Column) error {
	if err := kanbandb.SaveColumn(J.db.Columns, column); err != nil {
		return fmt.Errorf("failed saving column: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveColumns(columns []*types.Column) error {
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

func (J *JSONFile) MoveColumn(board, uuid string, toIndex int) error {
	b, err := kanbandb.GetBoard(J.db.Boards, board)
	if err != nil {
		return err
	}
	if err := kanbandb.MoveColumn(b, uuid, toIndex); err != nil {
		return err
	}
	return J.Save()
}
