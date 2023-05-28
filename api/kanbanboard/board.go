package kanbanboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/api"
	"sharkedule/kanban"
	"sharkedule/kanban/KTypes/description"
)

func Get(c *fiber.Ctx) error {
	board, err := kanban.ExtractBoard(c)
	if err != nil {
		return fmt.Errorf("failed to get board: %v", err)
	}
	return c.Status(fiber.StatusOK).JSON(board)
}

func List(c *fiber.Ctx) error {
	boards, err := kanban.List()
	if err != nil {
		return fmt.Errorf("failed getting boards: %v", err)
	}
	if err := c.Status(fiber.StatusOK).JSON(boards); err != nil {
		return fmt.Errorf("failed sending board: %v", err)

	}
	return nil
}

func ListNames(c *fiber.Ctx) error {

	boardNames, err := kanban.ListNames()
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

	var board BoardName

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&board); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending board: %v", err)
		}
	}

	b := kanban.New(board.Name)
	b.Description = &description.Description{
		Description: board.Description,
	}

	if err := b.Save(); err != nil {
		return err
	}

	if err := c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": b.UUID}); err != nil {
		return fmt.Errorf("failed sending board: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	b, err := kanban.ExtractBoard(c)
	if err != nil {
		return fmt.Errorf("failed to get board: %v", err)
	}
	if err := b.Delete(); err != nil {
		return fmt.Errorf("failed to delete board: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
