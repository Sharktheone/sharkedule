package access

import (
	"errors"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/column"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/tag"
	"github.com/Sharktheone/sharkedule/kanban/task"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/workspace"
	"slices"
)

func (a *Access) GetWorkspace(uuid string) (*workspace.Workspace, error) {
	if slices.Contains(a.Workspaces, uuid) {
		return db.DB.GetWorkspace(uuid)
	}
	return nil, errors.New("Workspace not found")
}

// Board functions
func (a *Access) CreateBoard(workspace, name string) (*board.Board, error) {

}

func (a *Access) SaveBoard(workspace string, board *board.Board) error {

}

func (a *Access) SaveBoards(workspace string, boards []*board.Board) error {

}

func (a *Access) GetBoard(workspace, uuid string) (*board.Board, error) {

}

func (a *Access) GetAllBoards(workspace string) ([]*board.Board, error) {

}

func (a *Access) GetBoards(workspace string, uuids []string) ([]*board.Board, error) {

}

func (a *Access) GetAllBoardNames(workspace string) ([]*namelist.NameList, error) {

}

func (a *Access) GetBoardNames(workspace string, uuids []string) (names []*namelist.NameList, err error) {
}

func (a *Access) DeleteBoard(workspace, uuid string) error {

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
