package kanban

import (
	"github.com/gofiber/fiber/v2"
	"sharkedule/database/db"
)

func GetBoard(uuid string) (*Board, error) {
	return db.DB.GetBoard(uuid)
}

func GetBoards() ([]*Board, error) {
	return db.DB.GetBoards()
}

func ExtractBoard(c *fiber.Ctx) (*Board, error) {
	boardUUID := c.Params("kanbanboard")
	return GetBoard(boardUUID)
}
