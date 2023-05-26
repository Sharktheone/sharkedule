package col

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/kanban"
	"sharkedule/kanban/KTypes"
)

func GetColumn(board interface{}, columnUUID string) (*KTypes.Column, int, error) {
	var b *KTypes.Board
	switch board := board.(type) {
	case *KTypes.Board:
		b = board
	case string:
		var err error
		b, _, err = kanban.GetBoard(board)
		if err != nil {
			return nil, -1, fmt.Errorf("failed getting board: %v", err)
		}

	}
	for i, column := range b.Columns {
		if column.UUID == columnUUID {
			return &column, i, nil
		}
	}

	return nil, -1, fmt.Errorf("column not found")
}

func ExtractColumn(c *fiber.Ctx) (*KTypes.Board, int, *KTypes.Column, int, error) {
	board, boardIndex, err := kanban.ExtractBoard(c)
	if err != nil {
		return nil, -1, nil, -1, fmt.Errorf("failed extracting board: %v", err)
	}

	columnUUID := c.Params("column")

	column, index, err := GetColumn(board, columnUUID)

	return board, boardIndex, column, index, nil
}
