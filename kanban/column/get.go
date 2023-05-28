package column

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/kanban"
)

func ExtractColumn(c *fiber.Ctx) (*kanban.Board, *Column, error) {
	board, err := kanban.ExtractBoard(c)
	if err != nil {
		return nil, nil, fmt.Errorf("failed extracting board: %v", err)
	}

	column, err := board.GetColumn(c.Params("column"))
	if err != nil {
		return board, nil, fmt.Errorf("failed getting column: %v", err)
	}

	return board, column, nil
}

func (c *Column) GetParentBoard() (*kanban.Board, error) {
	return kanban.GetBoard(c.Board)
}
