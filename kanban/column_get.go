package kanban

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ExtractColumn(c *fiber.Ctx) (*Board, *Column, error) {
	board, err := ExtractBoard(c)
	if err != nil {
		return nil, nil, fmt.Errorf("failed extracting board: %v", err)
	}

	column, err := board.GetColumn(c.Params("column"))
	if err != nil {
		return board, nil, fmt.Errorf("failed getting column: %v", err)
	}

	return board, column, nil
}

func (c *Column) GetParentBoard() (*Board, error) {
	return GetBoard(c.Board)
}
