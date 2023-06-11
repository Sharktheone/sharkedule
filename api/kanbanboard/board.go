package kanbanboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/sharkedule/api"
	"github.com/Sharktheone/sharkedule/api/middleware"
	board2 "github.com/Sharktheone/sharkedule/kanban/board"
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	b, err := middleware.ExtractBoard(c)
	if err != nil {
		return fmt.Errorf("failed to get board: %v", err)
	}
	return c.Status(fiber.StatusOK).JSON(b.Env())
}

func List(c *fiber.Ctx) error {
	boards, err := board2.GetBoards()
	if err != nil {
		return fmt.Errorf("failed getting boards: %v", err)
	}
	if err := c.Status(fiber.StatusOK).JSON(boards); err != nil {
		return fmt.Errorf("failed sending board: %v", err)

	}
	return nil
}

func ListNames(c *fiber.Ctx) error {
	boardNames, err := board2.Names()
	if err != nil {
		return fmt.Errorf("failed getting board names: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(boardNames)
}

func Create(c *fiber.Ctx) error {
	type BoardName struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var props BoardName

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&props); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending board: %v", err)
		}
	}

	b, err := board2.NewBoard(props.Name)
	if err != nil {
		return fmt.Errorf("failed creating board: %v", err)
	}
	b.Description = props.Description

	if err := b.Save(); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	if err := c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": b.UUID}); err != nil {
		return fmt.Errorf("failed sending board: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	b, err := middleware.ExtractBoard(c)
	if err != nil {
		return fmt.Errorf("failed to get board: %v", err)
	}
	if err := b.Delete(); err != nil {
		return fmt.Errorf("failed to delete board: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
