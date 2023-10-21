package jsonfile

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) CreateBoard(name string) (*types.Board, error) {
	b := kanbandb.CreateBoard(&J.db.Boards, name)
	return b, J.Save()
}

func (J *JSONFile) GetBoard(uuid string) (*types.Board, error) {
	return kanbandb.GetBoard(J.db.Boards, uuid)
}

func (J *JSONFile) GetAllBoards() ([]*types.Board, error) {
	return kanbandb.GetAllBoards(J.db.Boards), nil
}

func (J *JSONFile) GetBoards(uuids []string) ([]*types.Board, error) {
	return kanbandb.GetBoards(J.db.Boards, uuids)
}

func (J *JSONFile) GetAllBoardNames() ([]*namelist.NameList, error) {
	return kanbandb.GetAllBoardNames(J.db.Boards), nil
}

func (J *JSONFile) GetBoardNames(uuids []string) ([]*namelist.NameList, error) {
	return kanbandb.GetBoardNames(J.db.Boards, uuids)
}

func (J *JSONFile) SaveBoard(b *types.Board) error {
	if err := kanbandb.SaveBoard(J.db.Boards, b); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveBoards(boards []*types.Board) error {
	kanbandb.SaveBoards(J.db.Boards, boards)
	return J.Save()
}

func (J *JSONFile) DeleteBoard(uuid string) error {
	if err := kanbandb.DeleteBoard(J.db.Boards, uuid); err != nil {
		return err
	}
	return J.Save()
}
