package kanbanboard

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
	"sharkedule/api"
	"sharkedule/kanbanboardTypes"
)

var (
	KanbanBoard []kanbanboardTypes.KanbanBoard
)

func init() {
	loadTestBoard()
}

func loadTestBoard() {
	boards, err := os.Open("test_data.json")
	if err != nil {
		if os.IsNotExist(err) {
			KanbanBoard = []kanbanboardTypes.KanbanBoard{}
			log.Println("No test_data.json found, skipping loading test data")
			return
		}
		log.Fatalf("Error opening test_data.json: %v", err)
	}

	var boardsData []byte
	_, err = boards.Read(boardsData)
	if err != nil {
		log.Fatalf("Error reading test_data.json: %v", err)
	}

	if err := json.NewDecoder(boards).Decode(&KanbanBoard); err != nil {
		log.Fatalf("Error decoding test_data.json: %v", err)
	}
}

func getBoard(uuid string) (kanbanboardTypes.KanbanBoard, error) {
	if KanbanBoard == nil {
		loadTestBoard()
	}

	for _, board := range KanbanBoard {
		if board.UUID == uuid {
			return board, nil
		}
	}

	return kanbanboardTypes.KanbanBoard{}, errors.New("board not found")
}

func GetKanbanBoard(c *fiber.Ctx) error {
	boardUUID := c.Params("kanbanboard")

	if board, err := getBoard(boardUUID); err != nil || board.UUID == "" {
		errJson := api.JSON{"error": err.Error()}
		if sendErr := c.Status(http.StatusNotFound).JSON(errJson); sendErr != nil {
			log.Printf("Failed sending error (%v): %v", err, sendErr)
		}
	} else {
		if err := c.Status(http.StatusOK).JSON(board); err != nil {
			return fmt.Errorf("failed sending board: %v", err)
		}
	}

	return nil
}

func ListKanbanBoards(c *fiber.Ctx) error {
	if KanbanBoard == nil {
		loadTestBoard()
	}

	if err := c.Status(http.StatusOK).JSON(KanbanBoard); err != nil {
		return fmt.Errorf("failed sending board: %v", err)

	}
	return nil
}

func ListKanbanBoardNames(c *fiber.Ctx) error {
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

	if err := c.Status(http.StatusOK).JSON(boardNames); err != nil {
		return fmt.Errorf("failed sending board names: %v", err)

	}
	return nil
}

func CreateKanbanBoard(c *fiber.Ctx) error {
	type BoardName struct {
		Name string `json:"name"`
	}

	var board BoardName

	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&board); err != nil {
		if err := c.Status(http.StatusBadRequest).JSON(api.JSON{"error": err.Error()}); err != nil {
			return fmt.Errorf("failed sending board: %v", err)
		}
	}

	boardUUID := uuid.NewV4().String()

	var kBoard kanbanboardTypes.KanbanBoard

	kBoard.Name = board.Name
	kBoard.UUID = boardUUID

	KanbanBoard = append(KanbanBoard, kBoard)

	if err := c.Status(http.StatusOK).JSON(api.JSON{"uuid": boardUUID}); err != nil {
		return fmt.Errorf("failed sending board: %v", err)
	}
	return nil
}

func DeleteKanbanBoard(c *fiber.Ctx) error {
	boardUUID := c.Params("uuid")

	if board, err := getBoard(boardUUID); err != nil || board.UUID == "" {
		errJson := api.JSON{"error": err.Error()}
		if sendErr := c.Status(http.StatusNotFound).JSON(errJson); err != nil {
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
