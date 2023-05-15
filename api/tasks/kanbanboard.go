package tasks

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sharkedule/kanbanboard"
)

var (
	kanbanBoard []kanbanboard.KanbanBoard
)

func loadTestBoard() {
	boards, err := os.Open("test_data.json")
	if err != nil {
		log.Fatalf("Error opening test_data.json: %v", err)
	}

	var boardsData []byte
	_, err = boards.Read(boardsData)
	if err != nil {
		log.Fatalf("Error reading test_data.json: %v", err)
	}

	if err := json.NewDecoder(boards).Decode(&kanbanBoard); err != nil {
		log.Fatalf("Error decoding test_data.json: %v", err)
	}

}

func GetKanbanBoard(c *gin.Context) {

	uuid := c.Param("uuid")

	if board, err := getBoard(uuid); err != nil || board.UUID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, board)
		return
	}
}

func getBoard(uuid string) (kanbanboard.KanbanBoard, error) {
	if kanbanBoard == nil {
		loadTestBoard()
	}

	for _, board := range kanbanBoard {
		if board.UUID == uuid {
			return board, nil
		}
	}

	return kanbanboard.KanbanBoard{}, errors.New("board not found")
}
