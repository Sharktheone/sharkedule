package tasks

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

func ListKanbanBoards(c *gin.Context) {
	if kanbanBoard == nil {
		loadTestBoard()
	}

	c.JSON(http.StatusOK, kanbanBoard)
	return
}

func ListKanbanBoardNames(c *gin.Context) {
	if kanbanBoard == nil {
		loadTestBoard()
	}

	type BoardName struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}

	var boardNames []BoardName

	for _, board := range kanbanBoard {
		boardNames = append(boardNames, BoardName{UUID: board.UUID, Name: board.Name})
	}

	c.JSON(http.StatusOK, boardNames)
	return
}

func CreateKanbanBoard(c *gin.Context) {

	type BoardName struct {
		Name string `json:"name"`
	}

	var board BoardName

	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boardUUID := uuid.NewV4().String()

	var kBoard kanbanboard.KanbanBoard

	kBoard.Name = board.Name
	kBoard.UUID = boardUUID

	kanbanBoard = append(kanbanBoard, kBoard)

	c.JSON(http.StatusOK, gin.H{"uuid": boardUUID})
	return
}
