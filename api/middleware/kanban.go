package middleware

import (
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/column"
	"github.com/Sharktheone/sharkedule/kanban/task"
	"github.com/gofiber/fiber/v2"
)

func ExtractTask(c *fiber.Ctx) (*task.Task, error) {
	workspace := c.Params("workspace")
	uuid := c.Params("task")
	return task.Get(workspace, uuid)
}

func ExtractColumn(c *fiber.Ctx) (*column.Column, error) {
	workspace := c.Params("workspace")
	uuid := c.Params("column")
	return column.Get(workspace, uuid)
}

func ExtractBoard(c *fiber.Ctx) (*board.Board, error) {
	workspace := c.Params("workspace")
	boardUUID := c.Params("kanbanboard")
	return board.Get(workspace, boardUUID)
}
