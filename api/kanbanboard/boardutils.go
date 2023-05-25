package kanbanboard

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"sharkedule/kanbanboardTypes"
)

var (
	KanbanBoard []*kanbanboardTypes.KanbanBoard
)

func init() {
	loadTestBoard()
}

func loadTestBoard() {
	boards, err := os.Open("test_data.json")
	if err != nil {
		if os.IsNotExist(err) {
			KanbanBoard = []*kanbanboardTypes.KanbanBoard{}
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

func getBoard(uuid string) (*kanbanboardTypes.KanbanBoard, error) {
	if KanbanBoard == nil {
		loadTestBoard()
	}

	for _, board := range KanbanBoard {
		if board.UUID == uuid {
			return board, nil
		}
	}

	return &kanbanboardTypes.KanbanBoard{}, errors.New("board not found")
}

func extractBoard(c *fiber.Ctx) (*kanbanboardTypes.KanbanBoard, error) {
	boardUUID := c.Params("kanbanboard")

	board, err := getBoard(boardUUID)
	if err != nil {
		return &kanbanboardTypes.KanbanBoard{}, err
	}

	return board, nil
}
