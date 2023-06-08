package jsonfileV2

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
	kanbandb "github.com/Sharktheone/sharkedule/kanban/v2/database"
)

func (J *JSONFile) CreateBoard(name string) (error, *board.Board) {
	b := kanbandb.CreateBoard(J.db.Boards, name)
	return J.Save(), b
}

func (J *JSONFile) GetBoard(uuid string) (*board.Board, error) {
	return kanbandb.GetBoard(J.db.Boards, uuid)
}

func (J *JSONFile) GetBoards() ([]*board.Board, error) {
	return kanbandb.GetBoards(J.db.Boards), nil
}

func (J *JSONFile) GetBoardNames() ([]*namelist.NameList, error) {
	return kanbandb.GetBoardNames(J.db.Boards), nil
}

func (J *JSONFile) SaveBoard(b *board.Board) error {
	if err := kanbandb.SaveBoard(J.db.Boards, b); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveBoards(boards []*board.Board) error {
	kanbandb.SaveBoards(J.db.Boards, boards)
	return J.Save()
}

func (J *JSONFile) DeleteBoard(uuid string) error {
	if err := kanbandb.DeleteBoard(J.db.Boards, uuid); err != nil {
		return err
	}
	return J.Save()
}
