package database

import (
	"sharkedule/database/types"
	"sharkedule/kanban/KTypes/namelist"
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
	SaveColumn(boardUUID string, column *types.Column) error
	SaveColumns(boardUUID string, columns []*types.Column) error
	SaveTask(boardUUID string, column, task *types.Task) error
}

type DBStructure struct {
	Mu           *sync.Mutex
	Kanbanboards []*types.Board `json:"kanbanboards"`
}
