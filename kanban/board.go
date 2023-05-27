package kanban

import (
	"github.com/gofiber/fiber/v2"
	"sharkedule/database/db"
	"sharkedule/kanban/KTypes"
)

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
