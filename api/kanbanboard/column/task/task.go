package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"sharkedule/api"
	"sharkedule/database/db"
	"sharkedule/kanban"
	"sharkedule/kanban/KTypes"
	"sharkedule/kanban/col"
	"sharkedule/kanban/col/tsk"
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

	taskUUID := uuid.New().String()

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
	task := tsk.ExtractTask(c)
	if task.Err != nil {
		return fmt.Errorf("failed extracting task: %v", task.Err)
	}

	return c.Status(fiber.StatusOK).JSON(task.Task)
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

	task := tsk.ExtractTask(c)

	task.Column.Tasks = append(task.Column.Tasks[:task.ColumnIndex], task.Column.Tasks[task.ColumnIndex+1:]...)

	if moveTask.ToColumn == task.Column.UUID {
		task.Column.Tasks = append(task.Column.Tasks[:moveTask.ToIndex], append([]KTypes.KanbanTaskType{*task.Task}, task.Column.Tasks[moveTask.ToIndex:]...)...)
		task.Board.Columns[task.ColumnIndex] = *task.Column
		if err := db.DB.SaveBoard(task.Board); err != nil {
			return fmt.Errorf("failed saving board: %v", err)
		}
		return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task moved"})
	}

	task.Board.Columns[task.ColumnIndex] = *task.Column
	if err := db.DB.SaveBoard(task.Board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	toColumn, toIndex, err := col.GetColumn(task.Board, moveTask.ToColumn)
	if err != nil {
		return fmt.Errorf("failed getting column: %v", err)
	}
	toColumn.Tasks = append(toColumn.Tasks[:moveTask.ToIndex], append([]KTypes.KanbanTaskType{*task.Task}, toColumn.Tasks[moveTask.ToIndex:]...)...)
	task.Board.Columns[toIndex] = *toColumn

	if err := db.DB.SaveBoard(task.Board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task moved"})
}

func Delete(c *fiber.Ctx) error {

	task := tsk.ExtractTask(c)

	task.Column.Tasks = append(task.Column.Tasks[:task.TaskIndex], task.Column.Tasks[task.TaskIndex+1:]...)
	task.Board.Columns[task.ColumnIndex] = *task.Column

	if err := db.DB.SaveBoard(task.Board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task deleted"})
}
