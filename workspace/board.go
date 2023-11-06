package workspace

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (w *Workspace) CreateBoard(name string) (*board.Board, error) {
	b, err := db.DB.CreateBoard(w.UUID, name)
	if err != nil {
		return nil, err
	}
	return &board.Board{
		Board:     b,
		Workspace: w.UUID,
	}, nil
}

func (w *Workspace) SaveBoard(board *board.Board) error {
	return db.DB.SaveBoard(w.UUID, board.Board)
}

func (w *Workspace) SaveBoards(boards []*board.Board) error {
	var b []*types.Board

	for _, brd := range boards {
		b = append(b, brd.Board)
	}

	return db.DB.SaveBoards(w.UUID, b)
}

func (w *Workspace) GetBoard(uuid string) (*board.Board, error) {
	return board.Get(w.UUID, uuid)
}

func (w *Workspace) GetAllBoards() ([]*board.Board, error) {
	return board.GetBoardsAll(w.UUID)
}

func (w *Workspace) GetBoards(uuids []string) ([]*board.Board, error) {
	return board.GetBoards(w.UUID, uuids)
}

func (w *Workspace) GetAllBoardNames() ([]*namelist.NameList, error) {
	return board.AllNames(w.UUID)
}

func (w *Workspace) GetBoardNames(uuids []string) (names []*namelist.NameList, err error) {
	return board.Names(w.UUID, uuids)
}

func (w *Workspace) DeleteBoard(uuid string) error {
	return db.DB.DeleteBoard(w.UUID, uuid)
}
