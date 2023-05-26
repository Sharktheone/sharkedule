package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"sharkedule/api"
	"sharkedule/kanban"
	"sharkedule/kanban/KTypes"
)

func Create(c *fiber.Ctx) error {
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

	var task KTypes.KanbanTaskType

	task.Name = taskName.Name
	task.UUID = taskUUID

	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")

	for bIndex, board := range kanban.KBoard {
		if board.UUID == boardUUID {
			for index, column := range board.Columns {
				if column.UUID == columnUUID {
					column.Tasks = append(column.Tasks, task)
					board.Columns[index] = column
					kanban.KBoard[bIndex] = board

					return c.Status(fiber.StatusOK).JSON(task)

				}
			}

			return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "column not found"})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "board not found"})
}

func Get(c *fiber.Ctx) error {
	// TODO
	return nil
}

func Move(c *fiber.Ctx) error {
	type MoveTask struct {
		ToIndex  int    `json:"index"`
		ToColumn string `json:"column"`
	}

	var moveTask MoveTask

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&moveTask); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")
	taskUUID := c.Params("task")

	for bIndex, board := range kanban.KBoard {
		if board.UUID == boardUUID {
			for cIndex, column := range board.Columns {
				if column.UUID == columnUUID {
					for index, task := range column.Tasks {
						if task.UUID == taskUUID {
							column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)

							if moveTask.ToColumn == columnUUID {
								column.Tasks = append(column.Tasks[:moveTask.ToIndex], append([]KTypes.KanbanTaskType{task}, column.Tasks[moveTask.ToIndex:]...)...)
								board.Columns[cIndex] = column
								kanban.KBoard[bIndex] = board
								return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task moved"})
							}

							board.Columns[cIndex] = column
							kanban.KBoard[bIndex] = board

							for toIndex, toColumn := range board.Columns {
								if toColumn.UUID == moveTask.ToColumn {
									toColumn.Tasks = append(toColumn.Tasks[:moveTask.ToIndex], append([]KTypes.KanbanTaskType{task}, toColumn.Tasks[moveTask.ToIndex:]...)...)
									board.Columns[toIndex] = toColumn
									kanban.KBoard[bIndex] = board
									return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task moved"})
								}
							}
							return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "to_column not found"})
						}
					}
					return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "task not found"})
				}
			}
			return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "column not found"})
		}
	}

	return nil
}

func Delete(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")
	taskUUID := c.Params("task")

	for bIndex, board := range kanban.KBoard {
		if board.UUID == boardUUID {
			for cIndex, column := range board.Columns {
				if column.UUID == columnUUID {
					for index, task := range column.Tasks {
						if task.UUID == taskUUID {
							column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
							board.Columns[cIndex] = column
							kanban.KBoard[bIndex] = board
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
