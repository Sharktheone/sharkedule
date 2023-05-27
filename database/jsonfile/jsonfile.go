package jsonfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path"
	"sharkedule/database"
	"sharkedule/kanban/KTypes"
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
			Kanbanboards: []*KTypes.Board{},
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

	if err := json.NewDecoder(J.file).Decode(&J.db.Kanbanboards); err != nil {
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

func (J *JSONFile) SaveBoard(board *KTypes.Board) error {
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

func (J *JSONFile) SaveBoards(boards []*KTypes.Board) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	for _, board := range boards {
		if J.boardExists(board.UUID) {
			board, i, err := J.getBoard(board.UUID)
			if err != nil {
				return err
			}
			J.db.Kanbanboards[i] = board
		} else {
			return fmt.Errorf("board with uuid %s does not exist", board.UUID)
		}
	}
	if err := J.Save(); err != nil {
		return fmt.Errorf("failed saving database file: %v", err)
	}
	return nil
}

func (J *JSONFile) CreateBoard(boardName interface{}) error {
	var board *KTypes.Board
	switch b := boardName.(type) {
	case string:
		board = &KTypes.Board{
			Name: b,
		}
	case *KTypes.Board:
		board = b
	}
	board.UUID = uuid.New().String()

	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()

	J.db.Kanbanboards = append(J.db.Kanbanboards, board)
	return nil
}

func (J *JSONFile) GetBoard(boardUUID string) (*KTypes.Board, int, error) {
	board, index, err := J.getBoard(boardUUID)
	return board, index, err
}

func (J *JSONFile) GetBoards() ([]*KTypes.Board, error) {

	return J.db.Kanbanboards, nil
}

func (J *JSONFile) GetBoardNames() ([]*KTypes.NameList, error) {
	var boardNames []*KTypes.NameList
	for _, board := range J.db.Kanbanboards {
		boardName := &KTypes.NameList{
			Name: board.Name,
			UUID: board.UUID,
		}
		boardNames = append(boardNames, boardName)
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

func (J *JSONFile) getBoard(uuid string) (*KTypes.Board, int, error) {
	for i, board := range J.db.Kanbanboards {
		if board.UUID == uuid {
			return board, i, nil
		}
	}
	return &KTypes.Board{}, -1, database.ErrBoardNotFound
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

func (J *JSONFile) LockMutex() {
	J.db.Mu.Lock()
}

func (J *JSONFile) UnlockMutex() {
	J.db.Mu.Unlock()
}
