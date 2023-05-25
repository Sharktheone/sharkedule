package jsonfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sharkedule/database"
	"sharkedule/kanbanboardTypes"
)

const (
	DBFileName = "database.json"
)

type JSONFile struct {
	db   *database.DBStructure
	file *os.File
}

func (J *JSONFile) Load() error {
	dbPath := path.Join(database.DBRoot, DBFileName)
	file, err := os.Open(dbPath)
	if err != nil {
		return fmt.Errorf("failed opening database file: %v", err)
	}
	J.file = file
	defer J.file.Close()

	if err := json.NewDecoder(J.file).Decode(&J.db); err != nil {
		return fmt.Errorf("failed decoding database file: %v", err)
	}

	return nil
}

func (J *JSONFile) Save() error {
	var fileBuffer bytes.Buffer
	if err := json.NewEncoder(&fileBuffer).Encode(J.db); err != nil {
		return fmt.Errorf("failed encoding database file: %v", err)
	}
	if _, err := J.file.Write(fileBuffer.Bytes()); err != nil {
		return fmt.Errorf("failed writing database file: %v", err)
	}

	return nil
}

func (J *JSONFile) SaveBoard(board *kanbanboardTypes.KanbanBoard) error {
	if J.boardExists(board.UUID) {
		board, i, err := J.getBoard(board.UUID)
		if err != nil {
			return err
		}
		J.db.Kanbanboards[i] = *board

		if err := J.Save(); err != nil {
			return fmt.Errorf("failed saving database file: %v", err)
		}

	} else {
		return fmt.Errorf("board with uuid %s does not exist", board.UUID)
	}
	return nil
}

func (J *JSONFile) CreateBoard(boardName string) error {
	return nil
}

func (J *JSONFile) GetBoard(boardUUID string) (*kanbanboardTypes.KanbanBoard, error) {
	return nil, nil
}

func (J *JSONFile) GetBoards() ([]*kanbanboardTypes.KanbanBoard, error) {
	return nil, nil
}

func (J *JSONFile) GetBoardNames() ([]string, error) {
	return nil, nil
}

func (J *JSONFile) boardExists(uuid string) bool {
	for _, board := range J.db.Kanbanboards {
		if board.UUID == uuid {
			return true
		}
	}

	return false
}

func (J *JSONFile) getBoard(uuid string) (*kanbanboardTypes.KanbanBoard, int, error) {
	for i, board := range J.db.Kanbanboards {
		if board.UUID == uuid {
			return &board, i, nil
		}
	}
	return &kanbanboardTypes.KanbanBoard{}, -1, database.ErrBoardNotFound
}

func (J *JSONFile) writeToDisk() error {
	var fileBuffer bytes.Buffer
	if err := json.NewEncoder(&fileBuffer).Encode(J.db); err != nil {
		return fmt.Errorf("failed encoding database file: %v", err)
	}
	if _, err := J.file.Write(fileBuffer.Bytes()); err != nil {
		return fmt.Errorf("failed writing database file: %v", err)
	}
	return nil
}
