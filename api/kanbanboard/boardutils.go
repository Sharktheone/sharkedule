package kanbanboard

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"sharkedule/kanban"
)

var (
	KanbanBoard []*kanban.Board
)

func init() {
	loadTestBoard()
}

func loadTestBoard() {
	boards, err := os.Open("test_data.json")
	if err != nil {
		if os.IsNotExist(err) {
			KanbanBoard = []*kanban.Board{}
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

func getBoard(uuid string) (*kanban.Board, error) {
	if KanbanBoard == nil {
		loadTestBoard()
	}

	for _, board := range KanbanBoard {
		if board.UUID == uuid {
			return board, nil
		}
	}

	return &kanban.Board{}, errors.New("board not found")
}

func extractBoard(c *fiber.Ctx) (*kanban.Board, error) {
	boardUUID := c.Params("kanbanboard")

	board, err := getBoard(boardUUID)
	if err != nil {
		return &kanban.Board{}, err
	}

	return board, nil
}
