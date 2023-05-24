package database

import "sharkedule/kanbanboardTypes"

type IDatabase interface {
	Load() error
	Save() error
	SaveBoard(board *kanbanboardTypes.KanbanBoard) error
	CreateBoard(boardName string) error
	GetBoard(boardUUID string) (*kanbanboardTypes.KanbanBoard, error)
	GetBoards() ([]*kanbanboardTypes.KanbanBoard, error)
	GetBoardNames() ([]string, error)
}
