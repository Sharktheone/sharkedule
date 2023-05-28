package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/api"
	"sharkedule/kanban/column"
	"sharkedule/kanban/column/task"
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

	_, co, err := column.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("[CreateTask] failed extracting column: %v", err)
	}

	t := co.NewTask(taskName.Name)

	return c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": t.UUID})
}

func Get(c *fiber.Ctx) error {
	t, err := task.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(t)
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

	t, err := task.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	if err := t.Move(moveTask.ToIndex, moveTask.ToColumn); err != nil {
		return fmt.Errorf("failed moving task: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func Delete(c *fiber.Ctx) error {

	t, err := task.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)

	}
	if err := t.Delete(); err != nil {
		return fmt.Errorf("failed deleting task: %v", err)
	}
	return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task deleted"})
}
