package database

import (
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"sync"
)

const (
	DBRoot = "./"
)

type DBStructure struct {
	Mu *sync.Mutex `json:"-" yaml:"-" bson:"-"`
	types.Environment
}

type IDatabase interface {
	// Load and Save
	Load() error
	Save() error

	//Board functions
	CreateBoard(name string) (*types.Board, error)
	SaveBoard(board *types.Board) error
	SaveBoards(boards []*types.Board) error
	GetBoard(uuid string) (*types.Board, error)
	GetAllBoards() ([]*types.Board, error)
	GetBoards(uuids []string) ([]*types.Board, error)
	GetAllBoardNames() ([]*namelist.NameList, error)
	GetBoardNames(uuids []string) (names []*namelist.NameList, err error)
	DeleteBoard(uuid string) error

	//Column functions
	SaveColumn(column *types.Column) error
	SaveColumns(columns []*types.Column) error
	GetColumn(uuid string) (*types.Column, error)
	DeleteColumnOnBoard(board, column string) error
	RenameColumn(column, name string) error
	DeleteColumn(uuid string) error
	MoveColumn(board, uuid string, toIndex int) error
	NewColumn(board, name string) (*types.Column, error)

	//Task functions
	SaveTask(task *types.Task) error
	SaveTasks(tasks []*types.Task) error
	GetTask(uuid string) (*types.Task, error)
	DeleteTaskOnColumn(column, uuid string) error
	DeleteTask(uuid string) error
	MoveTask(column, uuid, toColumn string, toIndex int) error
	NewTask(column, name string) (*types.Task, error)
	RenameTask(task, name string) error
	RemoveTagOnTask(column, uuid string) error
	SetTagsOnTask(task string, tags []string) error
	SetTaskDescription(task, description string) error

	//Tag functions
	GetAllTags() ([]*types.Tag, error)
	GetTag(uuid string) (*types.Tag, error)
	AddTagToTask(task, tag string) error

	//Other functions
	GetStatus(uuid string) (*types.Status, error)
	GetPriority(uuid string) (*types.Priority, error)
	GetMember(uuid string) (*types.Member, error)
	GetChecklist(uuid string) (*types.Checklist, error)
	GetAttachment(uuid string) (*types.Attachment, error)
	GetDate(uuid string) (*types.Date, error)
}
