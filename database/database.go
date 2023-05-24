package database

import (
	"sharkedule/kanbanboardTypes"
	"sync"
)

const (
	DBRoot = "/var/lib/sharkedule"
)

type IDatabase interface {
	Load() error
	Save() error
	SaveBoard(board *kanbanboardTypes.KanbanBoard) error
	CreateBoard(boardName string) error
	GetBoard(boardUUID string) (*kanbanboardTypes.KanbanBoard, error)
	GetBoards() ([]*kanbanboardTypes.KanbanBoard, error)
	GetBoardNames() ([]string, error)
}

type DBStructure struct {
	Mu           *sync.Mutex
	Kanbanboards []kanbanboardTypes.KanbanBoard `json:"kanbanboards"`
}
