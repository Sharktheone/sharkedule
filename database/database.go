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
	CreateBoard(boardName string) error
	GetBoard(boardUUID string) (*KTypes.Board, error)
	GetBoards() ([]*KTypes.Board, error)
	GetBoardNames() ([]string, error)
}

type DBStructure struct {
	Mu           *sync.Mutex
	Kanbanboards []*KTypes.Board `json:"kanbanboards"`
}
