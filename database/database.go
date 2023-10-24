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
	Load() error
	Save() error
	SaveBoard(board *types.Board) error
	SaveBoards(boards []*types.Board) error
	SaveColumn(column *types.Column) error
	SaveColumns(columns []*types.Column) error
	SaveTask(task *types.Task) error
	SaveTasks(tasks []*types.Task) error
	CreateBoard(name string) (*types.Board, error)
	GetBoard(uuid string) (*types.Board, error)
	GetAllBoards() ([]*types.Board, error)
	GetBoards(uuids []string) ([]*types.Board, error)
	GetAllBoardNames() ([]*namelist.NameList, error)
	GetBoardNames(uuids []string) (names []*namelist.NameList, err error)
	GetColumn(uuid string) (*types.Column, error)
	GetTask(uuid string) (*types.Task, error)
	GetAllTags() ([]*types.Tag, error)
	GetTag(uuid string) (*types.Tag, error)
	GetStatus(uuid string) (*types.Status, error)
	GetPriority(uuid string) (*types.Priority, error)
	GetMember(uuid string) (*types.Member, error)
	GetChecklist(uuid string) (*types.Checklist, error)
	GetAttachment(uuid string) (*types.Attachment, error)
	GetDate(uuid string) (*types.Date, error)
	DeleteBoard(uuid string) error
	DeleteColumn(uuid string) error
	MoveColumn(board, uuid string, toIndex int) error
	DeleteTask(uuid string) error
	MoveTask(column, uuid, toColumn string, toIndex int) error
	DeleteTaskOnColumn(column, uuid string) error
	NewTask(column, name string) (*types.Task, error)
	NewColumn(board, name string) (*types.Column, error)
	AddTagToTask(task, tag string) error
	RemoveTagOnTask(column, uuid string) error
	SetTagsOnTask(task string, tags []string) error
	SetTaskDescription(task, description string) error
	RenameTask(task, name string) error
	DeleteColumnOnBoard(board, column string) error
	RenameColumn(column, name string) error
}
