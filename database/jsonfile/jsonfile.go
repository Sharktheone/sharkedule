package jsonfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sharkedule/database"
	"sharkedule/kanban"
	"sync"
)

const (
	DBFileName = "database.json"
)

type JSONFile struct {
	db   *database.DBStructure
	file *os.File
}

func NewJSONFile() *JSONFile {
	return &JSONFile{
		db: &database.DBStructure{
			Mu:           &sync.Mutex{},
			Kanbanboards: []*kanban.Board{},
		},
	}
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

func (J *JSONFile) SaveBoard(board *kanban.Board) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	if J.boardExists(board.UUID) {
		board, i, err := J.getBoard(board.UUID)
		if err != nil {
			return err
		}
		J.db.Kanbanboards[i] = board

		if err := J.Save(); err != nil {
			return fmt.Errorf("failed saving database file: %v", err)
		}

	} else {
		return fmt.Errorf("board with uuid %s does not exist", board.UUID)
	}
	return nil
}

func (J *JSONFile) CreateBoard(boardName string) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	board := &kanban.Board{
		Name: boardName,
	}
	J.db.Kanbanboards = append(J.db.Kanbanboards, board)
	return nil
}

func (J *JSONFile) GetBoard(boardUUID string) (*kanban.Board, error) {
	board, _, err := J.getBoard(boardUUID)
	return board, err
}

func (J *JSONFile) GetBoards() ([]*kanban.Board, error) {

	return J.db.Kanbanboards, nil
}

func (J *JSONFile) GetBoardNames() ([]string, error) {
	var boardNames []string
	for _, board := range J.db.Kanbanboards {
		boardNames = append(boardNames, board.Name)
	}
	return boardNames, nil
}

func (J *JSONFile) boardExists(uuid string) bool {
	for _, board := range J.db.Kanbanboards {
		if board.UUID == uuid {
			return true
		}
	}
	return false
}

func (J *JSONFile) getBoard(uuid string) (*kanban.Board, int, error) {
	for i, board := range J.db.Kanbanboards {
		if board.UUID == uuid {
			return board, i, nil
		}
	}
	return &kanban.Board{}, -1, database.ErrBoardNotFound
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
