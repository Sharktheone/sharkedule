package jsonfile

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path"
	"sharkedule/database"
	"sharkedule/database/types"
	"sharkedule/kanban/KTypes/namelist"
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
			Kanbanboards: []*types.Board{},
		},
	}
}

func (J *JSONFile) Load() error {
	dbPath := path.Join(database.DBRoot, DBFileName)
	file, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("failed opening database file: %v", err)
	}
	J.file = file

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

func (J *JSONFile) SaveBoard(board *types.Board) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	if J.boardExists(board.UUID) {
		board, err := J.getBoard(board.UUID)
		if err != nil {
			return err
		}
		J.db.Kanbanboards[board.Index] = board

		if err := J.Save(); err != nil {
			return fmt.Errorf("failed saving database file: %v", err)
		}

	} else {
		return fmt.Errorf("board with uuid %s does not exist", board.UUID)
	}
	return nil
}

func (J *JSONFile) SaveBoards(boards []*types.Board) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	for _, board := range boards {
		if J.boardExists(board.UUID) {
			board, err := J.getBoard(board.UUID)
			if err != nil {
				return err
			}
			J.db.Kanbanboards[board.Index] = board
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
	var board *types.Board
	switch b := boardName.(type) {
	case string:
		board = &types.Board{
			Name: b,
		}
	case *types.Board:
		board = b
	}
	board.UUID = uuid.New().String()

	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()

	J.db.Kanbanboards = append(J.db.Kanbanboards, board)
	return nil
}

func (J *JSONFile) GetBoard(boardUUID string) (*types.Board, error) {
	return J.getBoard(boardUUID)
}

func (J *JSONFile) GetBoards() ([]*types.Board, error) {

	return J.db.Kanbanboards, nil
}

func (J *JSONFile) GetBoardNames() ([]*namelist.NameList, error) {
	var boardNames []*namelist.NameList
	for _, board := range J.db.Kanbanboards {
		boardName := &namelist.NameList{
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

func (J *JSONFile) getBoard(uuid string) (*types.Board, error) {
	for _, board := range J.db.Kanbanboards {
		if board.UUID == uuid {
			attachValues(board)
			return board, nil
		}
	}
	return &types.Board{}, database.ErrBoardNotFound
}

func attachValues(board *types.Board) {
	for colIndex, column := range board.Columns {
		column.Board = board.UUID
		column.Index = colIndex

		for taskIndex, task := range column.Tasks {
			task.Board = board.UUID
			task.Column = column.UUID
			task.Index = taskIndex
		}
	}
}

func attachMultiple(boards []*types.Board) {
	for _, board := range boards {
		attachValues(board)
	}
}

func (J *JSONFile) writeToDisk() error {
	var fileBuffer bytes.Buffer
	attachMultiple(J.db.Kanbanboards)
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

func (J *JSONFile) GetColumn(boardUUID string, columnUUID string) (*types.Column, error) {
	if !J.boardExists(boardUUID) {
		return &types.Column{}, fmt.Errorf("board with uuid %s does not exist", boardUUID)
	}
	board, err := J.getBoard(boardUUID)
	if err != nil {
		return &types.Column{}, err
	}
	for _, column := range board.Columns {
		if column.UUID == columnUUID {
			return column, nil
		}
	}
	return &types.Column{}, errors.New("column not found")
}

func (J *JSONFile) SaveColumn(boardUUID string, column *types.Column) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	if !J.boardExists(boardUUID) {
		return fmt.Errorf("board with uuid %s does not exist", boardUUID)
	}
	board, err := J.getBoard(boardUUID)
	if err != nil {
		return err
	}
	if board.Columns[column.Index].UUID != column.UUID {
		for index, col := range board.Columns {
			if col.UUID == column.UUID {
				column.Index = index
				break
			}
		}
	}
	board.Columns[column.Index] = column

	if err := J.Save(); err != nil {
		return fmt.Errorf("failed saving database file: %v", err)
	}
	return nil
}

func (J *JSONFile) SaveColumns(boardUUID string, columns []*types.Column) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	if !J.boardExists(boardUUID) {
		return fmt.Errorf("board with uuid %s does not exist", boardUUID)
	}
	board, err := J.getBoard(boardUUID)
	if err != nil {
		return err
	}
	for _, column := range columns {
		if board.Columns[column.Index].UUID != column.UUID {
			for index, col := range board.Columns {
				if col.UUID == column.UUID {
					column.Index = index
					break
				}
			}
		}
		board.Columns[column.Index] = column
	}

	if err := J.Save(); err != nil {
		return fmt.Errorf("failed saving database file: %v", err)
	}
	return nil
}

func (J *JSONFile) GetTask(boardUUID string, columnUUID string, taskUUID string) (*types.Task, error) {
	if !J.boardExists(boardUUID) {
		return &types.Task{}, fmt.Errorf("board with uuid %s does not exist", boardUUID)
	}
	board, err := J.getBoard(boardUUID)
	if err != nil {
		return &types.Task{}, err
	}
	for _, column := range board.Columns {
		if column.UUID == columnUUID {
			for _, task := range column.Tasks {
				if task.UUID == taskUUID {
					return task, nil
				}
			}
		}
	}
	return &types.Task{}, errors.New("task not found")
}

func (J *JSONFile) SaveTask(boardUUID string, column, task *types.Task) error {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	if !J.boardExists(boardUUID) {
		return fmt.Errorf("board with uuid %s does not exist", boardUUID)
	}
	board, err := J.getBoard(boardUUID)
	if err != nil {
		return err
	}
	if board.Columns[column.Index].UUID != column.UUID {
		for index, col := range board.Columns {
			if col.UUID == column.UUID {
				column.Index = index
				break
			}
		}
	}
	if board.Columns[column.Index].Tasks[task.Index].UUID != task.UUID {
		for index, tsk := range board.Columns[column.Index].Tasks {
			if tsk.UUID == task.UUID {
				task.Index = index
				break
			}
		}
	}

	board.Columns[column.Index].Tasks[task.Index] = task

	if err := J.Save(); err != nil {
		return fmt.Errorf("failed saving database file: %v", err)
	}
	return nil
}
