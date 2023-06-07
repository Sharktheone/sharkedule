package jsonfileV2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/sharkedule/database"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
	kanbandb "github.com/Sharktheone/sharkedule/kanban/v2/database"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
	"os"
	"path"
	"sync"
)

const (
	DBFileName = "database.json"
)

type JSONFile struct {
	db   *database.DBStructureV2
	path string
}

func NewJSONFile() *JSONFile {
	return &JSONFile{
		db: &database.DBStructureV2{
			Mu: &sync.Mutex{},
		},
	}
}

func (J *JSONFile) Load() error {
	dbPath := path.Join(database.DBRoot, DBFileName)
	file, err := os.OpenFile(dbPath, os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("failed opening database file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("failed closing database file:", err)
		}
	}()

	J.path = dbPath

	if err := json.NewDecoder(file).Decode(&J.db); err != nil {
		return fmt.Errorf("failed decoding database file: %v", err)
	}

	return nil
}

func (J *JSONFile) Save() error {
	var fileBuffer bytes.Buffer
	var encoder = json.NewEncoder(&fileBuffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(J.db); err != nil {
		return fmt.Errorf("failed encoding database file: %v", err)
	}

	if err := os.WriteFile(J.path, fileBuffer.Bytes(), 0755); err != nil {
		return fmt.Errorf("failed writing database file: %v", err)
	}

	return nil
}

func (J *JSONFile) SaveBoard(board *types2.Board) error {
	if err := kanbandb.SaveBoard(J.db.Boards, board); err != nil {
		return fmt.Errorf("failed saving board: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveBoards(boards []*types2.Board) error {
	kanbandb.SaveBoards(J.db.Boards, boards)
	return J.Save()
}

func (J *JSONFile) SaveColumn(column *types2.Column) error {
	if err := kanbandb.SaveColumn(J.db.Columns, column); err != nil {
		return fmt.Errorf("failed saving column: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveColumns(columns []*types2.Column) error {
	kanbandb.SaveColumns(J.db.Columns, columns)
	return J.Save()
}

func (J *JSONFile) SaveTask(task *task.Task) error {
	if err := kanbandb.SaveTask(J.db.Tasks, task); err != nil {
		return fmt.Errorf("failed saving task: %v", err)
	}
	return J.Save()
}

func (J *JSONFile) SaveTasks(tasks []*task.Task) error {
	kanbandb.SaveTasks(J.db.Tasks, tasks)
	return J.Save()
}

func (J *JSONFile) CreateBoard(name string) (error, *types2.Board) {
	board := kanbandb.CreateBoard(J.db.Boards, name)
	return J.Save(), board
}

func (J *JSONFile) GetBoard(uuid string) (*types2.Board, error) {
	return kanbandb.GetBoard(J.db.Boards, uuid)
}

func (J *JSONFile) GetBoards() ([]*types2.Board, error) {
	return kanbandb.GetBoards(J.db.Boards), nil
}

func (J *JSONFile) GetBoardNames() ([]*namelist.NameList, error) {
	return kanbandb.GetBoardNames(J.db.Boards), nil
}

func (J *JSONFile) GetColumn(uuid string) (*types2.Column, error) {
	return kanbandb.GetColumn(J.db.Columns, uuid)
}

func (J *JSONFile) GetTask(uuid string) (*task.Task, error) {
	return kanbandb.GetTask(J.db.Tasks, uuid)
}

func (J *JSONFile) GetTag(uuid string) (*types2.Tag, error) {
	return kanbandb.GetTag(J.db.Tags, uuid)
}

func (J *JSONFile) GetTags() ([]*types2.Tag, error) {
	return kanbandb.GetTags(J.db.Tags), nil
}

func (J *JSONFile) GetStatus(uuid string) (*types2.Status, error) {
	return kanbandb.GetStatus(J.db.Status, uuid)
}

func (J *JSONFile) GetPriority(uuid string) (*types2.Priority, error) {
	return kanbandb.GetPriority(J.db.Priority, uuid)
}

func (J *JSONFile) GetMember(uuid string) (*types2.Member, error) {
	return kanbandb.GetMember(J.db.Members, uuid)
}

func (J *JSONFile) GetChecklist(uuid string) (*types2.Checklist, error) {
	return kanbandb.GetChecklist(J.db.Checklists, uuid)
}

func (J *JSONFile) GetAttachment(uuid string) (*types2.Attachment, error) {
	return kanbandb.GetAttachment(J.db.Attachments, uuid)
}

func (J *JSONFile) GetDate(uuid string) (*types2.Date, error) {
	return kanbandb.GetDate(J.db.Dates, uuid)
}
