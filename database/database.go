package database

import (
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
	"sync"
)

const (
	DBRoot = "./"
)

type DBStructure struct {
	Mu *sync.Mutex `json:"-" yaml:"-" bson:"-"`
	types2.Environment
}

type IDatabase interface {
	Load() error
	Save() error
	SaveBoard(board *types2.Board) error
	SaveBoards(boards []*types2.Board) error
	SaveColumn(column *types2.Column) error
	SaveColumns(columns []*types2.Column) error
	SaveTask(task *types2.Task) error
	SaveTasks(tasks []*types2.Task) error
	CreateBoard(name string) (*types2.Board, error)
	GetBoard(uuid string) (*types2.Board, error)
	GetAllBoards() ([]*types2.Board, error)
	GetBoardNames() ([]*namelist.NameList, error)
	GetColumn(uuid string) (*types2.Column, error)
	GetTask(uuid string) (*types2.Task, error)
	GetTags() ([]*types2.Tag, error)
	GetTag(uuid string) (*types2.Tag, error)
	GetStatus(uuid string) (*types2.Status, error)
	GetPriority(uuid string) (*types2.Priority, error)
	GetMember(uuid string) (*types2.Member, error)
	GetChecklist(uuid string) (*types2.Checklist, error)
	GetAttachment(uuid string) (*types2.Attachment, error)
	GetDate(uuid string) (*types2.Date, error)
	DeleteBoard(uuid string) error
	DeleteColumn(uuid string) error
	MoveColumn(board, uuid string, toIndex int) error
	DeleteTask(uuid string) error
	MoveTask(column, uuid, toColumn string, toIndex int) error
	DeleteTaskOnColumn(column, uuid string) error
	NewTask(column, name string) (*types2.Task, error)
	NewColumn(board, name string) (*types2.Column, error)
	AddTagToTask(task, tag string) error
	RemoveTagOnTask(column, uuid string) error
	SetTagsOnTask(task string, tags []string) error
	SetTaskDescription(task, description string) error
	RenameTask(task, name string) error
	DeleteColumnOnBoard(board, column string) error
	RenameColumn(column, name string) error
}
