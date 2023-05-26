package kanban

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"sharkedule/kanban/KTypes"
)

var (
	KBoard []*KTypes.Board
)

func init() {
	LoadTestBoard()
}

func LoadTestBoard() {
	boards, err := os.Open("test_data.json")
	if err != nil {
		if os.IsNotExist(err) {
			KBoard = []*KTypes.Board{}
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

	if err := json.NewDecoder(boards).Decode(&KBoard); err != nil {
		log.Fatalf("Error decoding test_data.json: %v", err)
	}
}

func GetBoard(uuid string) (*KTypes.Board, error) {
	if KBoard == nil {
		LoadTestBoard()
	}

	for _, board := range KBoard {
		if board.UUID == uuid {
			return board, nil
		}
	}

	return &KTypes.Board{}, errors.New("board not found")
}

func ExtractBoard(c *fiber.Ctx) (*KTypes.Board, error) {
	boardUUID := c.Params("kanbanboard")

	board, err := GetBoard(boardUUID)
	if err != nil {
		return &KTypes.Board{}, err
	}

	return board, nil
}
