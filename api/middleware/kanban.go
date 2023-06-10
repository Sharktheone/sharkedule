package middleware

import (
	board2 "github.com/Sharktheone/sharkedule/kanban/board"
	column2 "github.com/Sharktheone/sharkedule/kanban/column"
	task2 "github.com/Sharktheone/sharkedule/kanban/task"
	"github.com/gofiber/fiber/v2"
)

func ExtractTask(c *fiber.Ctx) (*task2.Task, error) {
	uuid := c.Params("task")
	return task2.Get(uuid)
}

func ExtractColumn(c *fiber.Ctx) (*column2.Column, error) {
	uuid := c.Params("column")
	return column2.Get(uuid)
}

func ExtractBoard(c *fiber.Ctx) (*board2.Board, error) {
	boardUUID := c.Params("kanbanboard")
	return board2.Get(boardUUID)
}
