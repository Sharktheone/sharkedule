package kanbanboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"sharkedule/api"
	"sharkedule/database/db"
	"sharkedule/kanban"
	"sharkedule/kanban/KTypes"
)

func Get(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")

	if board, _, err := kanban.GetBoard(boardUUID); err != nil || board.UUID == "" {
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
	boards, err := db.DB.GetBoards()
	if err != nil {
		return fmt.Errorf("failed getting boards: %v", err)
	}
	if err := c.Status(fiber.StatusOK).JSON(boards); err != nil {
		return fmt.Errorf("failed sending board: %v", err)

	}
	return nil
}

func ListNames(c *fiber.Ctx) error {
	type BoardName struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}

	boardNames, err := db.DB.GetBoardNames()
	if err != nil {
		return fmt.Errorf("failed getting board names: %v", err)
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

	boardUUID := uuid.New().String()

	kBoard := &KTypes.Board{
		Name: board.Name,
		Description: KTypes.Description{
			Description: board.Description,
		},
		UUID: boardUUID,
	}

	if err := db.DB.CreateBoard(kBoard); err != nil {
		return fmt.Errorf("failed creating board: %v", err)
	}

	if err := c.Status(fiber.StatusOK).JSON(api.JSON{"uuid": boardUUID}); err != nil {
		return fmt.Errorf("failed sending board: %v", err)
	}
	return nil
}

func Delete(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")

	db.DB.LockMutex()
	defer db.DB.UnlockMutex()

	boards, err := db.DB.GetBoards()
	if err != nil {
		return fmt.Errorf("failed getting boards: %v", err)
	}

	if board, _, err := kanban.GetBoard(boardUUID); err != nil || board.UUID == "" {
		errJson := api.JSON{"error": err.Error()}
		if sendErr := c.Status(fiber.StatusBadRequest).JSON(errJson); err != nil {
			log.Printf("Failed sending error (%v): %v", err, sendErr)
		}
	} else {
		for i, board := range boards {
			if board.UUID == boardUUID {
				boards = append(boards[:i], boards[i+1:]...)
				if err := db.DB.SaveBoards(boards); err != nil {
					return fmt.Errorf("failed setting boards: %v", err)
				}
				break
			}
		}
	}

	return nil
}
