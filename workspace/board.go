package workspace

import (
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
)

func (w *Workspace) CreateBoard(workspace, name string) (*board.Board, error) {

}

func (w *Workspace) SaveBoard(workspace string, board *board.Board) error {

}

func (w *Workspace) SaveBoards(workspace string, boards []*board.Board) error {

}

func (w *Workspace) GetBoard(workspace, uuid string) (*board.Board, error) {

}

func (w *Workspace) GetAllBoards(workspace string) ([]*board.Board, error) {

}

func (w *Workspace) GetBoards(workspace string, uuids []string) ([]*board.Board, error) {

}

func (w *Workspace) GetAllBoardNames(workspace string) ([]*namelist.NameList, error) {

}

func (w *Workspace) GetBoardNames(workspace string, uuids []string) (names []*namelist.NameList, err error) {
}

func (w *Workspace) DeleteBoard(workspace, uuid string) error {

}
