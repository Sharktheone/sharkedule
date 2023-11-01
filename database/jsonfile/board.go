package jsonfile

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) CreateBoard(name string) (*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	b := kanbandb.CreateBoard(&ws.Boards, name)
	return b, J.Save()
}

func (J *JSONFile) GetBoard(uuid string) (*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetBoard(ws.Boards, uuid)
}

func (J *JSONFile) GetAllBoards() ([]*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetAllBoards(ws.Boards), nil
}

func (J *JSONFile) GetBoards(uuids []string) ([]*types.Board, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetBoards(ws.Boards, uuids)
}

func (J *JSONFile) GetAllBoardNames() ([]*namelist.NameList, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetAllBoardNames(ws.Boards), nil
}

func (J *JSONFile) GetBoardNames(uuids []string) ([]*namelist.NameList, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetBoardNames(ws.Boards, uuids)
}

func (J *JSONFile) SaveBoard(b *types.Board) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.SaveBoard(ws.Boards, b); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveBoards(boards []*types.Board) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	kanbandb.SaveBoards(ws.Boards, boards)
	return J.Save()
}

func (J *JSONFile) DeleteBoard(uuid string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.DeleteBoard(ws.Boards, uuid); err != nil {
		return err
	}
	return J.Save()
}
