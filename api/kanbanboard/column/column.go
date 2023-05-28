package column

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/api"
	"sharkedule/database/db"
	"sharkedule/kanban/column"
	"sharkedule/kanban/old/col"
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

	board, _, co, _, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board.Columns = append(board.Columns, co)
	if err := db.DB.SaveBoard(board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "unknown error"})
}

func Move(c *fiber.Ctx) error {
	type MoveColumn struct {
		UUID  string `json:"uuid"`
		Index int    `json:"index"`
	}

	var moveColumn MoveColumn

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&moveColumn); err != nil {
		if err := c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending task: %v", err)
		}
	}

	board, _, co, index, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)

	board.Columns = append(board.Columns[:moveColumn.Index], append([]*column.Column{co}, board.Columns[moveColumn.Index:]...)...)
	if err := db.DB.SaveBoard(board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "unknown error"})
}

func Get(c *fiber.Ctx) error {
	_, _, co, _, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := c.Status(fiber.StatusOK).JSON(co); err != nil {
		return fmt.Errorf("failed sending column: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	board, _, _, index, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)
	if err := db.DB.SaveBoard(board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
