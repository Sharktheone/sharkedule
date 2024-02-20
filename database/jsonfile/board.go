package jsonfile

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
	types2 "github.com/Sharktheone/sharkedule/types"
)

func (J *JSONFile) CreateBoard(workspace, name string) (*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	b := kanbandb.CreateBoard(&ws.Boards, name)
	return b, J.Save()
}

func (J *JSONFile) GetBoard(workspace, uuid string) (*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetBoard(ws.Boards, uuid)
}

func (J *JSONFile) GetAllBoards(workspace string) ([]*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetAllBoards(ws.Boards), nil
}

func (J *JSONFile) GetBoards(workspace string, uuids []string) ([]*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetBoards(ws.Boards, uuids)
}

func (J *JSONFile) GetAllBoardNames(workspace string) ([]*types2.NameList, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetAllBoardNames(ws.Boards), nil
}

func (J *JSONFile) GetBoardNames(workspace string, uuids []string) ([]*types2.NameList, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetBoardNames(ws.Boards, uuids)
}

func (J *JSONFile) SaveBoard(workspace string, b *types.Board) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.SaveBoard(ws.Boards, b); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveBoards(workspace string, boards []*types.Board) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	kanbandb.SaveBoards(ws.Boards, boards)
	return J.Save()
}

func (J *JSONFile) DeleteBoard(workspace, uuid string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.DeleteBoard(ws.Boards, uuid); err != nil {
		return err
	}
	return J.Save()
}
