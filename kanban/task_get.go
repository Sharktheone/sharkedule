package kanban

import (
	"github.com/gofiber/fiber/v2"
)

func ExtractTask(c *fiber.Ctx) (*Task, error) {
	_, co, err := ExtractColumn(c)
	if err != nil {
		return nil, err
	}
	taskUUID := c.Params("task")
	return co.GetTask(taskUUID)

}

func (t *Task) GetParentBoard() (*Board, error) {
	return GetBoard(t.Board)
}

func (t *Task) GetParentColumn() (*Column, error) {
	board, err := t.GetParentBoard()
	if err != nil {
		return nil, err
	}
	return board.GetColumn(t.Column)
}
