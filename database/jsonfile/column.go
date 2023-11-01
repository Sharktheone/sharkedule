package jsonfile

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) NewColumn(board, name string) (*types.Column, error) {
	b, err := J.GetBoard(board)
	if err != nil {
		return nil, err
	}
	c := kanbandb.NewColumn(&J.db.Columns, b, name)
	return c, J.Save()
}

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

func (J *JSONFile) DeleteColumnOnBoard(board, column string) error {
	b, err := kanbandb.GetBoard(J.db.Boards, board)
	if err != nil {
		return err
	}
	if err := kanbandb.DeleteColumnOnBoard(b, column); err != nil {
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

func (J *JSONFile) RenameColumn(column, name string) error {
	col, err := J.GetColumn(column)
	if err != nil {
		return err
	}
	kanbandb.RenameColumn(col, name)

	return J.Save()
}
