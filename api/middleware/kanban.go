package middleware

import (
	"github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/Sharktheone/sharkedule/kanban/column"
	"github.com/Sharktheone/sharkedule/kanban/task"
	"github.com/Sharktheone/sharkedule/user"
	"github.com/gofiber/fiber/v2"
)

func ExtractTask(c *fiber.Ctx) (*user.User, *task.Task, error) {
	workspace := c.Params("workspace")
	uuid := c.Params("task")

	u, err := ExtractUser(c)
	if err != nil {
		return nil, nil, err
	}

	t, err := u.Access.GetTask(workspace, uuid)
	return u, t, err
}

func ExtractColumn(c *fiber.Ctx) (*user.User, *column.Column, error) {
	workspace := c.Params("workspace")
	uuid := c.Params("column")

	u, err := ExtractUser(c)
	if err != nil {
		return nil, nil, err
	}

	col, err := u.Access.GetColumn(workspace, uuid)

	return u, col, err
}

func ExtractBoard(c *fiber.Ctx) (*user.User, *board.Board, error) {
	workspace := c.Params("workspace")
	uuid := c.Params("board")

	u, err := ExtractUser(c)
	if err != nil {
		return nil, nil, err
	}

	brd, err := u.Access.GetBoard(workspace, uuid)

	return u, brd, err
}
