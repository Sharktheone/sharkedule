package middleware

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
	"github.com/gofiber/fiber/v2"
)

func ExtractTask(c *fiber.Ctx) (*task.Task, error) {
	uuid := c.Params("task")
	return task.Get(uuid)
}

func ExtractColumn(c *fiber.Ctx) (*column.Column, error) {
	uuid := c.Params("column")
	return column.Get(uuid)
}

func ExtractBoard(c *fiber.Ctx) (*board.Board, error) {
	boardUUID := c.Params("kanbanboard")
	return board.Get(boardUUID)
}
