package col

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/kanban"
	"sharkedule/kanban/column"
	"sharkedule/kanban/old"
)

func GetColumn(board interface{}, col interface{}) (*column.Column, int, error) {
	var b *kanban.Board
	switch board := board.(type) {
	case *kanban.Board:
		b = board
	case string:
		var err error
		b, _, err = old.GetBoard(board)
		if err != nil {
			return nil, -1, fmt.Errorf("failed getting board: %v", err)
		}
	default:
		return nil, -1, fmt.Errorf("invalid board type")
	}

	var columnUUID string

	switch c := col.(type) {
	case string:
		columnUUID = c
	case *column.Column:
		columnUUID = c.UUID
	default:
		return nil, -1, fmt.Errorf("invalid column type")

	}

	for i, c := range b.Columns {
		if c.UUID == columnUUID {
			return c, i, nil
		}
	}

	return nil, -1, fmt.Errorf("column not found")
}

func ExtractColumn(c *fiber.Ctx) (*kanban.Board, int, *column.Column, int, error) {
	board, boardIndex, err := old.ExtractBoard(c)
	if err != nil {
		return nil, -1, nil, -1, fmt.Errorf("failed extracting board: %v", err)
	}

	columnUUID := c.Params("column")

	co, index, err := GetColumn(board, columnUUID)

	return board, boardIndex, co, index, nil
}
