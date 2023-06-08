package jsonfileV2

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
	kanbandb "github.com/Sharktheone/sharkedule/kanban/v2/database"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func (J *JSONFile) CreateBoard(name string) (*types.Board, error) {
	b := kanbandb.CreateBoard(J.db.Boards, name)
	return b, J.Save()
}

func (J *JSONFile) GetBoard(uuid string) (*types.Board, error) {
	return kanbandb.GetBoard(J.db.Boards, uuid)
}

func (J *JSONFile) GetBoards() ([]*types.Board, error) {
	return kanbandb.GetBoards(J.db.Boards), nil
}

func (J *JSONFile) GetBoardNames() ([]*namelist.NameList, error) {
	return kanbandb.GetBoardNames(J.db.Boards), nil
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
