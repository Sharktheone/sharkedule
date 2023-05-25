package kanbanboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"log"
	"sharkedule/api"
	"sharkedule/kanban"
)

func Get(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")

	if board, err := getBoard(boardUUID); err != nil || board.UUID == "" {
		errJson := api.JSON{"error": err.Error()}
		if sendErr := c.Status(fiber.StatusNotFound).JSON(errJson); sendErr != nil {
			log.Printf("Failed sending error (%v): %v", err, sendErr)
		}
	} else {
		if err := c.Status(fiber.StatusOK).JSON(board); err != nil {
			return fmt.Errorf("failed sending board: %v", err)
		}
	}

	return nil
}

func List(c *fiber.Ctx) error {
	if KanbanBoard == nil {
		loadTestBoard()
	}

	if err := c.Status(fiber.StatusOK).JSON(KanbanBoard); err != nil {
		return fmt.Errorf("failed sending board: %v", err)

	}
	return nil
}

func ListNames(c *fiber.Ctx) error {
	if KanbanBoard == nil {
		loadTestBoard()
	}

	type BoardName struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}

	var boardNames []BoardName

	for _, board := range KanbanBoard {
		boardNames = append(boardNames, BoardName{UUID: board.UUID, Name: board.Name})
	}

	if err := c.Status(fiber.StatusOK).JSON(boardNames); err != nil {
		return fmt.Errorf("failed sending board names: %v", err)

	}
	return nil
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

	boardUUID := uuid.NewV4().String()

	var kBoard *kanban.Board

	kBoard.Name = board.Name
	kBoard.Description.Description = board.Description
	kBoard.UUID = boardUUID

	KanbanBoard = append(KanbanBoard, kBoard)

	if err := c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": boardUUID}); err != nil {
		return fmt.Errorf("failed sending board: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")

	if board, err := getBoard(boardUUID); err != nil || board.UUID == "" {
		errJson := api.JSON{"error": err.Error()}
		if sendErr := c.Status(fiber.StatusBadRequest).JSON(errJson); err != nil {
			log.Printf("Failed sending error (%v): %v", err, sendErr)
		}
	} else {
		for i, board := range KanbanBoard {
			if board.UUID == boardUUID {
				KanbanBoard = append(KanbanBoard[:i], KanbanBoard[i+1:]...)
				break
			}
		}
	}

	return nil
}
