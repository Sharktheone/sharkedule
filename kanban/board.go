package kanban

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sharkedule/database/db"
	"sharkedule/kanban/KTypes"
)

var (
	KBoard []*KTypes.Board // TODO: Replace all usages with db.DB.GetBoards()
)

func Load() {
	var err error
	KBoard, err = db.DB.GetBoards()
	if err != nil {
		log.Fatalf("failed getting boards: %v", err)
	}
}

func GetBoard(board interface{}) (*KTypes.Board, int, error) {
	var uuid string
	switch board := board.(type) {
	case string:
		uuid = board
	case *KTypes.Board:
		uuid = board.UUID
	}

	return db.DB.GetBoard(uuid)
}

func ExtractBoard(c *fiber.Ctx) (*KTypes.Board, int, error) {
	boardUUID := c.Params("kanbanboard")
	return GetBoard(boardUUID)
}
