package middlewear

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/board/column"
	"github.com/Sharktheone/sharkedule/kanban/board/column/task"
	"github.com/gofiber/fiber/v2"
)

func ExtractTask(c *fiber.Ctx) (*task.Task, error) {
	_, co, err := ExtractColumn(c)
	if err != nil {
		return nil, err
	}
	taskUUID := c.Params("task")
	return co.GetTask(taskUUID)
}

func ExtractColumn(c *fiber.Ctx) (*board.Board, *column.Column, error) {
	b, err := ExtractBoard(c)
	if err != nil {
		return nil, nil, fmt.Errorf("failed extracting board: %v", err)
	}

	column, err := b.GetColumn(c.Params("column"))
	if err != nil {
		return b, nil, fmt.Errorf("failed getting column: %v", err)
	}

	return b, column, nil
}

func ExtractBoard(c *fiber.Ctx) (*board.Board, error) {
	boardUUID := c.Params("kanbanboard")
	return board.GetBoard(boardUUID)
}
