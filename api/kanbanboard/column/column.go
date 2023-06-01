package column

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/api"
	"sharkedule/api/middlewear"
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

	b, err := middlewear.ExtractBoard(c)
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

	_, co, err := middlewear.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := co.Move(moveColumn.Index); err != nil {
		return fmt.Errorf("failed moving column: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func Get(c *fiber.Ctx) error {
	_, co, err := middlewear.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := c.Status(fiber.StatusOK).JSON(co); err != nil {
		return fmt.Errorf("failed sending column: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	_, co, err := middlewear.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := co.Delete(); err != nil {
		return fmt.Errorf("failed deleting column: %v", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
