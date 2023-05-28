package old

import (
	"github.com/gofiber/fiber/v2"
	"sharkedule/database/db"
	"sharkedule/kanban"
)

func GetBoard(board interface{}) (*kanban.Board, int, error) {
	var uuid string
	switch board := board.(type) {
	case string:
		uuid = board
	case *kanban.Board:
		uuid = board.UUID
	}
	b, err := db.DB.GetBoard(uuid)
	if err != nil {
		return nil, -1, err
	}

	return b, b.Index, err
}

func ExtractBoard(c *fiber.Ctx) (*kanban.Board, int, error) {
	boardUUID := c.Params("kanbanboard")
	return GetBoard(boardUUID)
}
