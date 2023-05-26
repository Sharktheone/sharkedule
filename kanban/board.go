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

func GetBoard(uuid string) (*KTypes.Board, error) {
	return db.DB.GetBoard(uuid)
}

func ExtractBoard(c *fiber.Ctx) (*KTypes.Board, error) {
	boardUUID := c.Params("kanbanboard")
	return GetBoard(boardUUID)
}
