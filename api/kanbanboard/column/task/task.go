package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/api"
	"sharkedule/database/db"
	"sharkedule/kanban/column/task"
	"sharkedule/kanban/old/col"
	"sharkedule/kanban/old/col/tsk"
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

	var t task.Task

	t.Name = taskName.Name

	boardUUID := c.Params("kanbanboard")
	columnUUID := c.Params("column")

	if uuid, err := tsk.Create(boardUUID, columnUUID, &t); err != nil {
		return fmt.Errorf("failed creating task: %v", err)
	} else {
		return c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": uuid})
	}
}

func Get(c *fiber.Ctx) error {
	t := tsk.ExtractTask(c)
	if t.Err != nil {
		return fmt.Errorf("failed extracting task: %v", t.Err)
	}

	return c.Status(fiber.StatusOK).JSON(t.Task)
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

	t := tsk.ExtractTask(c)

	t.Column.Tasks = append(t.Column.Tasks[:t.ColumnIndex], t.Column.Tasks[t.ColumnIndex+1:]...)

	if moveTask.ToColumn == t.Column.UUID {
		t.Column.Tasks = append(t.Column.Tasks[:moveTask.ToIndex], append([]*task.Task{t.Task}, t.Column.Tasks[moveTask.ToIndex:]...)...)
		t.Board.Columns[t.ColumnIndex] = t.Column
		if err := db.DB.SaveBoard(t.Board); err != nil {
			return fmt.Errorf("failed saving board: %v", err)
		}
		return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task moved"})
	}

	t.Board.Columns[t.ColumnIndex] = t.Column
	if err := db.DB.SaveBoard(t.Board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	toColumn, toIndex, err := col.GetColumn(t.Board, moveTask.ToColumn)
	if err != nil {
		return fmt.Errorf("failed getting column: %v", err)
	}
	toColumn.Tasks = append(toColumn.Tasks[:moveTask.ToIndex], append([]*task.Task{t.Task}, toColumn.Tasks[moveTask.ToIndex:]...)...)
	t.Board.Columns[toIndex] = toColumn

	if err := db.DB.SaveBoard(t.Board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task moved"})
}

func Delete(c *fiber.Ctx) error {

	t := tsk.ExtractTask(c)

	t.Column.Tasks = append(t.Column.Tasks[:t.TaskIndex], t.Column.Tasks[t.TaskIndex+1:]...)
	t.Board.Columns[t.ColumnIndex] = t.Column

	if err := db.DB.SaveBoard(t.Board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task deleted"})
}
