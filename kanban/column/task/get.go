package task

import (
	"github.com/gofiber/fiber/v2"
	"sharkedule/kanban"
	"sharkedule/kanban/column"
)

func ExtractTask(c *fiber.Ctx) (*Task, error) {
	_, co, err := column.ExtractColumn(c)
	if err != nil {
		return nil, err
	}
	taskUUID := c.Params("task")
	return co.GetTask(taskUUID)

}

func (t *Task) GetParentBoard() (*kanban.Board, error) {
	return kanban.GetBoard(t.Board)
}

func (t *Task) GetParentColumn() (*column.Column, error) {
	board, err := t.GetParentBoard()
	if err != nil {
		return nil, err
	}
	return board.GetColumn(t.Column)
}
