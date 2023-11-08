package access

import (
	"errors"
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/column"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/tag"
	"github.com/Sharktheone/sharkedule/kanban/task"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/workspace"
)

func (a *Access) GetWorkspace(uuid string) (*workspace.Workspace, error) {
	for _, w := range a.Workspaces {
		if w.UUID == uuid {
			return db.DB.GetWorkspace(uuid)
		}
	}
	return nil, errors.New("workspace not found")
}

// Board functions
func (a *Access) CreateBoard(workspace, name string) (*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.Permissions.CreateBoards {
		return nil, fmt.Errorf("no permissions to create board in workspace %s", workspace)
	}

	brd, err := db.DB.CreateBoard(workspace, name)
	if err != nil {
		return nil, err
	}
	return &board.Board{Board: brd, Workspace: workspace}, nil
}

func (a *Access) SaveBoard(workspace string, board *board.Board) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateBoards {
		return fmt.Errorf("no permissions to update board in workspace %s", workspace)
	}

	brd, err := ws.board(board.UUID)
	if err == nil {
		if !brd.Permissions.Update {
			return fmt.Errorf("no permissions to update board %s", board.UUID)
		}
	}

	return db.DB.SaveBoard(workspace, board.Board)

}

func (a *Access) SaveBoards(workspace string, boards []*board.Board) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.UpdateBoards {
		return fmt.Errorf("no permissions to update board in workspace %s", workspace)
	}

	for _, b := range boards {
		brd, err := ws.board(b.UUID)
		if err == nil { //error is nil, when we have
			if !brd.Permissions.Update {
				return fmt.Errorf("no permissions to update board %s", b.UUID)
			}
		}

	}

	var b []*types.Board
	for _, brd := range boards {
		b = append(b, brd.Board)
	}

	return db.DB.SaveBoards(workspace, b)

}

func (a *Access) GetBoard(workspace, uuid string) (*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		_, err := ws.board(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	b, err := db.DB.GetBoard(workspace, uuid)
	if err != nil {
		return nil, err
	}
	return &board.Board{Board: b, Workspace: workspace}, nil

}

func (a *Access) GetAllBoards(workspace string) ([]*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		var uuids []string
		for _, b := range ws.Boards {
			uuids = append(uuids, b.UUID)
		}
		b, err := db.DB.GetBoards(workspace, uuids)
		if err != nil {
			return nil, err
		}

		var boards []*board.Board
		for _, brd := range b {
			boards = append(boards, &board.Board{Board: brd, Workspace: workspace})
		}
		return boards, nil
	}

	b, err := db.DB.GetAllBoards(workspace)
	if err != nil {
		return nil, err
	}

	var boards []*board.Board
	for _, brd := range b {
		boards = append(boards, &board.Board{Board: brd, Workspace: workspace})
	}
	return boards, nil

}

func (a *Access) GetBoards(workspace string, uuids []string) ([]*board.Board, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		for _, uuid := range uuids {
			_, err := ws.board(uuid) //when the board is NOT in the slice / the user can't access them, this returns an error
			if err != nil {
				return nil, err
			}
		}
	}

	b, err := db.DB.GetBoards(workspace, uuids)
	if err != nil {
		return nil, err
	}

	var boards []*board.Board
	for _, brd := range b {
		boards = append(boards, &board.Board{Board: brd, Workspace: workspace})
	}
	return boards, nil

}

func (a *Access) GetAllBoardNames(workspace string) ([]*namelist.NameList, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		var uuids []string
		for _, b := range ws.Boards {
			uuids = append(uuids, b.UUID)
		}
		return db.DB.GetBoardNames(workspace, uuids)
	}

	return db.DB.GetAllBoardNames(workspace)

}

func (a *Access) GetBoardNames(workspace string, uuids []string) (names []*namelist.NameList, err error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllBoards {
		for _, uuid := range uuids {
			_, err := ws.board(uuid) //when the board is NOT in the slice / the user can't access them, this returns an error
			if err != nil {
				return nil, err
			}
		}
	}

	return db.DB.GetBoardNames(workspace, uuids)
}

func (a *Access) DeleteBoard(workspace, uuid string) error {
	ws, err := a.workspace(workspace)
	if err != nil {
		return err
	}

	if !ws.Permissions.DeleteBoards {
		return fmt.Errorf("no permissions to delete board in workspace %s", workspace)
	}

	brd, err := ws.board(uuid)
	if err == nil {
		if !brd.Permissions.Delete {
			return fmt.Errorf("no permissions to delete board %s", uuid)
		}
	}

	return db.DB.DeleteBoard(workspace, uuid)

}

// Column functions
func (a *Access) SaveColumn(workspace string, column *column.Column) error {

}

func (a *Access) SaveColumns(workspace string, columns []*column.Column) error {

}

func (a *Access) GetColumn(workspace, uuid string) (*column.Column, error) {

}

func (a *Access) DeleteColumnOnBoard(workspace, board, column string) error {

}

func (a *Access) RenameColumn(workspace, column, name string) error {

}

func (a *Access) DeleteColumn(workspace, uuid string) error {

}

func (a *Access) MoveColumn(workspace, board, uuid string, toIndex int) error {

}

func (a *Access) NewColumn(workspace, board, name string) (*column.Column, error) {

}

// Task functions
func (a *Access) SaveTask(workspace string, task *task.Task) error {

}

func (a *Access) SaveTasks(workspace string, tasks []*task.Task) error {

}

func (a *Access) GetTask(workspace, uuid string) (*task.Task, error) {

}

func (a *Access) DeleteTaskOnColumn(workspace, column, uuid string) error {

}

func (a *Access) DeleteTask(workspace, uuid string) error {

}

func (a *Access) MoveTask(workspace, column, uuid, toColumn string, toIndex int) error {

}

func (a *Access) NewTask(workspace, column, name string) (*task.Task, error) {

}

func (a *Access) RenameTask(workspace, task, name string) error {

}

func (a *Access) RemoveTagOnTask(workspace, column, uuid string) error {

}

func (a *Access) SetTagsOnTask(workspace, task string, tags []string) error {

}

func (a *Access) SetTaskDescription(workspace, task, description string) error {

}

// Tag functions
func (a *Access) GetAllTags(workspace string) ([]*tag.Tag, error) {

}

func (a *Access) GetTag(workspace, uuid string) (*tag.Tag, error) {

}

func (a *Access) AddTagToTask(workspace, task, tag string) error {

}

//Other functions

func (a *Access) GetStatus(workspace, uuid string) (*types.Status, error) {

}

func (a *Access) GetPriority(workspace, uuid string) (*types.Priority, error) {

}

func (a *Access) GetChecklist(workspace, uuid string) (*types.Checklist, error) {

}

func (a *Access) GetAttachment(workspace, uuid string) (*types.Attachment, error) {

}

func (a *Access) GetDate(workspace, uuid string) (*types.Date, error) {

}

//GetUser(uuid string) (*types.Member, error) TODO

func (a *Access) workspace(uuid string) (*WorkspaceAccess, error) {
	for _, w := range a.Workspaces {
		if w.UUID == uuid {
			return &w, nil
		}
	}
	return nil, errors.New("workspace not found")
}

//func (a *Access) board(workspace, uuid string) (*BoardAccess, error) {
//	ws, err := a.workspace(workspace)
//	if err != nil {
//		return nil, err
//	}
//
//	return ws.board(uuid)
//}
