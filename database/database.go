package database

import (
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/workspace"
	"sync"
)

const (
	DBRoot = "./"
)

type DBStructure struct {
	Mu         *sync.Mutex            `json:"-" yaml:"-" bson:"-"`
	Workspaces []*workspace.Workspace `json:"workspaces" yaml:"workspaces" bson:"workspaces"`
}

type IDatabase interface {
	// Load and Save
	Load() error
	Save() error

	//Board functions
	CreateBoard(workspace, name string) (*types.Board, error)
	SaveBoard(workspace string, board *types.Board) error
	SaveBoards(workspace string, boards []*types.Board) error
	GetBoard(workspace, uuid string) (*types.Board, error)
	GetAllBoards() ([]*types.Board, error)
	GetBoards(workspace string, uuids []string) ([]*types.Board, error)
	GetAllBoardNames(workspace string) ([]*namelist.NameList, error)
	GetBoardNames(workspace string, uuids []string) (names []*namelist.NameList, err error)
	DeleteBoard(workspace, uuid string) error

	//Column functions
	SaveColumn(workspace string, column *types.Column) error
	SaveColumns(workspace string, columns []*types.Column) error
	GetColumn(workspace, uuid string) (*types.Column, error)
	DeleteColumnOnBoard(workspace, board, column string) error
	RenameColumn(workspace, column, name string) error
	DeleteColumn(workspace, uuid string) error
	MoveColumn(workspace, board, uuid string, toIndex int) error
	NewColumn(workspace, board, name string) (*types.Column, error)

	//Task functions
	SaveTask(workspace string, task *types.Task) error
	SaveTasks(workspace string, tasks []*types.Task) error
	GetTask(workspace, uuid string) (*types.Task, error)
	DeleteTaskOnColumn(workspace, column, uuid string) error
	DeleteTask(workspace, uuid string) error
	MoveTask(workspace, column, uuid, toColumn string, toIndex int) error
	NewTask(workspace, column, name string) (*types.Task, error)
	RenameTask(workspace, task, name string) error
	RemoveTagOnTask(workspace, column, uuid string) error
	SetTagsOnTask(workspace, task string, tags []string) error
	SetTaskDescription(workspace, task, description string) error

	//Tag functions
	GetAllTags(workspace string) ([]*types.Tag, error)
	GetTag(workspace, uuid string) (*types.Tag, error)
	AddTagToTask(workspace, task, tag string) error

	//Other functions
	GetStatus(workspace, uuid string) (*types.Status, error)
	GetPriority(workspace, uuid string) (*types.Priority, error)
	GetMember(workspace, uuid string) (*types.Member, error)
	GetChecklist(workspace, uuid string) (*types.Checklist, error)
	GetAttachment(workspace, uuid string) (*types.Attachment, error)
	GetDate(workspace, uuid string) (*types.Date, error)
}
