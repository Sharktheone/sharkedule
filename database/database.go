package database

import (
	"github.com/Sharktheone/sharkedule/database/types"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
	"sync"
)

const (
	DBRoot = "/var/lib/sharkedule"
)

type IDatabase interface {
	Load() error
	Save() error
	SaveBoard(board *types.Board) error
	SaveBoards(boards []*types.Board) error
	CreateBoard(boardName interface{}) error
	GetBoard(boardUUID string) (*types.Board, error)
	GetBoards() ([]*types.Board, error)
	GetBoardNames() ([]*namelist.NameList, error)
	LockMutex()
	UnlockMutex()
	GetColumn(boardUUID, columnUUID string) (*types.Column, error)
	SaveColumn(boardUUID string, column *types.Column) error
	SaveColumns(boardUUID string, columns []*types.Column) error
	GetTask(boardUUID, columnUUID, taskUUID string) (*types.Task, error)
	SaveTask(boardUUID string, column, task *types.Task) error
}

type DBStructure struct {
	Mu           *sync.Mutex    `json:"-" yaml:"-" bson:"-"`
	Kanbanboards []*types.Board `json:"kanbanboards"`
}

type DBStructureV2 struct {
	Mu *sync.Mutex `json:"-" yaml:"-" bson:"-"`
	types2.Environment
}

type IDatabaseV2 interface {
	Load() error
	Save() error
	SaveBoard(board *types2.Board) error
	SaveBoards(boards []*types2.Board) error
	SaveColumn(column *types2.Column) error
	SaveColumns(columns []*types2.Column) error
	SaveTask(task *types2.Task) error
	SaveTasks(tasks []*types2.Task) error
	CreateBoard(name string) (error, *types2.Board)
	GetBoard(uuid string) (*types2.Board, error)
	GetBoards() ([]*types2.Board, error)
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
}
