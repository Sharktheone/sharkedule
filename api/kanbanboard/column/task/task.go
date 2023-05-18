package task

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

func CreateKanbanBoardColumnTask(c *fiber.Ctx) error {
	type CreateTask struct {
		Name string `json:"name"`
	}

	var taskName CreateTask

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&taskName); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	taskUUID := uuid.NewV4().String()

	var task kanbanboardTypes.KanbanTaskType

	task.Name = taskName.Name
	task.UUID = taskUUID

	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")

	for bIndex, board := range kanbanboard.KanbanBoard {
		if board.UUID == boardUUID {
			for index, column := range board.Columns {
				if column.UUID == columnUUID {
					column.Tasks = append(column.Tasks, task)
					board.Columns[index] = column
					kanbanboard.KanbanBoard[bIndex] = board

					return c.Status(fiber.StatusOK).JSON(task)

				}
			}

			return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "column not found"})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "board not found"})
}

func GetKanbanBoardColumnTask(c *fiber.Ctx) error {
	// TODO
	return nil
}

func DeleteKanbanBoardColumnTask(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")
	taskUUID := c.Params("task")

	for bIndex, board := range kanbanboard.KanbanBoard {
		if board.UUID == boardUUID {
			for cIndex, column := range board.Columns {
				if column.UUID == columnUUID {
					for index, task := range column.Tasks {
						if task.UUID == taskUUID {
							column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
							board.Columns[cIndex] = column
							kanbanboard.KanbanBoard[bIndex] = board
							return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task deleted"})
						}
					}
					return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "task not found"})
				}
			}
			return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "column not found"})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "board not found"})
}
