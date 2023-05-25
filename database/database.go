package database

import (
	"sharkedule/kanban"
	"sync"
)

const (
	DBRoot = "/var/lib/sharkedule"
)

type IDatabase interface {
	Load() error
	Save() error
	SaveBoard(board *kanban.Board) error
	CreateBoard(boardName string) error
	GetBoard(boardUUID string) (*kanban.Board, error)
	GetBoards() ([]*kanban.Board, error)
	GetBoardNames() ([]string, error)
}

type DBStructure struct {
	Mu           *sync.Mutex
	Kanbanboards []*kanban.Board `json:"kanbanboards"`
}
