package column

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"sharkedule/api"
	"sharkedule/api/kanbanboard"
	"sharkedule/kanbanboardTypes"
)

func CreateKanbanBoardColumn(c *fiber.Ctx) error {
	type CreateColumn struct {
		Name string `json:"name"`
	}

	var boardName CreateColumn

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&boardName); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	columnUUID := uuid.NewV4().String()

	var column kanbanboardTypes.KanbanColumnType

	column.Name = boardName.Name
	column.UUID = columnUUID

	boardUUID := c.Params("kanbanboard")

	for index, board := range kanbanboard.KanbanBoard {
		if board.UUID == boardUUID {
			board.Columns = append(board.Columns, column)
			kanbanboard.KanbanBoard[index] = board

			return c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": columnUUID})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "board not found"})
}

func GetKanbanBoardColumn(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")

	for _, board := range kanbanboard.KanbanBoard {
		if board.UUID == boardUUID {
			for _, column := range board.Columns {
				if column.UUID == columnUUID {
					if err := c.Status(fiber.StatusOK).JSON(column); err != nil {
						return fmt.Errorf("failed sending column: %v", err)
					}
				}
			}
		}
	}

	return nil
}

func DeleteKanbanBoardColumn(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")

	for bIndex, board := range kanbanboard.KanbanBoard {
		if board.UUID == boardUUID {
			for index, column := range board.Columns {
				if column.UUID == columnUUID {
					board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)
					kanbanboard.KanbanBoard[bIndex] = board
					return c.SendStatus(fiber.StatusOK)
				}
			}
			return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "column not found"})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "board not found"})
}
