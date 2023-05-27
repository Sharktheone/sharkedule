package database

import (
	"sharkedule/kanban/KTypes"
	"sync"
)

const (
	DBRoot = "/var/lib/sharkedule"
)

type IDatabase interface {
	Load() error
	Save() error
	SaveBoard(board *KTypes.Board) error
	SaveBoards(boards []*KTypes.Board) error
	CreateBoard(boardName interface{}) error
	GetBoard(boardUUID string) (*KTypes.Board, int, error)
	GetBoards() ([]*KTypes.Board, error)
	GetBoardNames() ([]*KTypes.NameList, error)
	LockMutex()
	UnlockMutex()
}

type DBStructure struct {
	Mu           *sync.Mutex
	Kanbanboards []*KTypes.Board `json:"kanbanboards"`
}
