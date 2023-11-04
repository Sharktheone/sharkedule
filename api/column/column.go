package column

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/sharkedule/api"
	"github.com/Sharktheone/sharkedule/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	type CreateColumn struct {
		Name string `json:"name"`
	}

	var boardName CreateColumn

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&boardName); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	b, err := middleware.ExtractBoard(c)
	if err != nil {
		return fmt.Errorf("failed extracting board: %v", err)
	}
	co, err := b.NewColumn(boardName.Name)
	if err != nil {
		return fmt.Errorf("failed creating column: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": co.UUID})
}

func Move(c *fiber.Ctx) error {
	type MoveColumn struct {
		Index int `json:"index"`
	}

	var moveColumn MoveColumn

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&moveColumn); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	co, err := middleware.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board := c.Params("board")

	if err := co.Move(board, moveColumn.Index); err != nil {
		return fmt.Errorf("failed moving column: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func Get(c *fiber.Ctx) error {
	co, err := middleware.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := c.Status(fiber.StatusOK).JSON(co.Env()); err != nil {
		return fmt.Errorf("failed sending column: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	co, err := middleware.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := co.Delete(); err != nil {
		return fmt.Errorf("failed deleting column: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func DeleteOnBoard(c *fiber.Ctx) error {
	co, err := middleware.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board := c.Params("board")

	if err := co.DeleteOnBoard(board); err != nil {
		return fmt.Errorf("failed deleting column: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func Rename(c *fiber.Ctx) error {
	type RenameColumn struct {
		Name string `json:"name"`
	}

	var renameColumn RenameColumn

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&renameColumn); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	co, err := middleware.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := co.Rename(renameColumn.Name); err != nil {
		return fmt.Errorf("failed renaming column: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
