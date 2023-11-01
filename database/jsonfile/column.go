package jsonfile

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) NewColumn(workspace, board, name string) (*types.Column, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	b, err := kanbandb.GetBoard(ws.Boards, board)
	if err != nil {
		return nil, err
	}
	c := kanbandb.NewColumn(&ws.Columns, b, name)
	return c, J.Save()
}

func (J *JSONFile) GetColumn(workspace, uuid string) (*types.Column, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetColumn(ws.Columns, uuid)
}

func (J *JSONFile) SaveColumn(workspace string, column *types.Column) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.SaveColumn(ws.Columns, column); err != nil {
		return fmt.Errorf("failed saving column: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveColumns(workspace string, columns []*types.Column) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	kanbandb.SaveColumns(ws.Columns, columns)
	return J.Save()
}

func (J *JSONFile) DeleteColumn(workspace string, uuid string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.DeleteColumn(ws.Columns, uuid); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) DeleteColumnOnBoard(workspace, board, column string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	b, err := kanbandb.GetBoard(ws.Boards, board)
	if err != nil {
		return err
	}
	if err := kanbandb.DeleteColumnOnBoard(b, column); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) MoveColumn(workspace, board, uuid string, toIndex int) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	b, err := kanbandb.GetBoard(ws.Boards, board)
	if err != nil {
		return err
	}
	if err := kanbandb.MoveColumn(b, uuid, toIndex); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) RenameColumn(workspace, column, name string) error {
	col, err := J.GetColumn(workspace, column)
	if err != nil {
		return err
	}
	kanbandb.RenameColumn(col, name)

	return J.Save()
}
