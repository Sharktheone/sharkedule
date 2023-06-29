package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/sharkedule/api"
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/gofiber/fiber/v2"
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

	co, err := middleware.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("[CreateTask] failed extracting column: %v", err)
	}

	t, err := db.DB.NewTask(co.Column.UUID, taskName.Name)
	if err != nil {
		return fmt.Errorf("failed creating task: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": t.UUID})
}

func Get(c *fiber.Ctx) error {
	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(t.Env())
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

	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	column := c.Params("column")

	if err := t.Move(column, moveTask.ToColumn, moveTask.ToIndex); err != nil {
		return fmt.Errorf("failed moving task: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func Delete(c *fiber.Ctx) error {
	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)

	}
	if err := t.Delete(); err != nil {
		return fmt.Errorf("failed deleting task: %v", err)
	}
	return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task deleted"})
}

func DeleteOnColumn(c *fiber.Ctx) error {
	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}
	column := c.Params("column")

	if err := t.DeleteOnColumn(column); err != nil {
		return fmt.Errorf("failed deleting task: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(api.JSON{"success": "task deleted"})
}

func Rename(c *fiber.Ctx) error {
	type RenameTask struct {
		Name string `json:"name"`
	}

	var renameTask RenameTask

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&renameTask); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	if err := t.Rename(renameTask.Name); err != nil {
		return fmt.Errorf("failed renaming task: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func AddTag(c *fiber.Ctx) error {
	type AddTag struct {
		Tag string `json:"tag"`
	}

	var addTag AddTag

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&addTag); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	if err := t.AddTag(addTag.Tag); err != nil {
		return fmt.Errorf("failed adding tags: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func RemoveTag(c *fiber.Ctx) error {
	type RemoveTag struct {
		Tag string `json:"tag"`
	}

	var removeTag RemoveTag

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&removeTag); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	if err := t.RemoveTag(removeTag.Tag); err != nil {
		return fmt.Errorf("failed removing tags: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func SetTags(c *fiber.Ctx) error {
	type SetTags struct {
		Tags []string `json:"tags"`
	}

	var setTags SetTags

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&setTags); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	t, err := middleware.ExtractTask(c)
	if err != nil {
		return fmt.Errorf("failed extracting task: %v", err)
	}

	if err := t.SetTags(setTags.Tags); err != nil {
		return fmt.Errorf("failed setting tags: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
