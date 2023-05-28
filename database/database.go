package database

import (
	"sharkedule/kanban"
	"sharkedule/kanban/KTypes/namelist"
	"sync"
)

const (
	DBRoot = "/var/lib/sharkedule"
)

type IDatabase interface {
	Load() error
	Save() error
	SaveBoard(board *kanban.Board) error
	SaveBoards(boards []*kanban.Board) error
	CreateBoard(boardName interface{}) error
	GetBoard(boardUUID string) (*kanban.Board, error)
	GetBoards() ([]*kanban.Board, error)
	GetBoardNames() ([]*namelist.NameList, error)
	LockMutex()
	UnlockMutex()
}

type DBStructure struct {
	Mu           *sync.Mutex
	Kanbanboards []*kanban.Board `json:"kanbanboards"`
}
