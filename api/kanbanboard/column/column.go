package column

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sharkedule/api"
	"sharkedule/database/db"
	"sharkedule/kanban/KTypes"
	col "sharkedule/kanban/column"
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

	board, column, _, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board.Columns = append(board.Columns, *column)
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

	board, column, index, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)

	board.Columns = append(board.Columns[:moveColumn.Index], append([]KTypes.Column{*column}, board.Columns[moveColumn.Index:]...)...)
	if err := db.DB.SaveBoard(board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}

	return c.Status(fiber.StatusBadRequest).JSON(api.JSON{"error": "unknown error"})
}

func Get(c *fiber.Ctx) error {
	_, column, _, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	if err := c.Status(fiber.StatusOK).JSON(column); err != nil {
		return fmt.Errorf("failed sending column: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	board, _, index, err := col.ExtractColumn(c)
	if err != nil {
		return fmt.Errorf("failed extracting column: %v", err)
	}

	board.Columns = append(board.Columns[:index], board.Columns[index+1:]...)
	if err := db.DB.SaveBoard(board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
