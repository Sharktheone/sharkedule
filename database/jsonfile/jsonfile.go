package jsonfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/sharkedule/database"
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"os"
	"path"
	"sync"
)

const (
	DBFileName = "database.json"
)

type JSONFile struct {
	db   *database.DBStructure
	path string
}

func NewJSONFile() *JSONFile {
	return &JSONFile{
		db: &database.DBStructure{
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

func (J *JSONFile) GetStatus(workspace, uuid string) (*types.Status, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetStatus(ws.Statuses, uuid)
}

func (J *JSONFile) GetPriority(workspace, uuid string) (*types.Priority, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetPriority(ws.Priorities, uuid)
}

//func (J *JSONFile) GetUser(uuid string) (*types.Member, error) {
//	return kanbandb.GetUser(J.db.Users, uuid) //TODO
//}

func (J *JSONFile) GetChecklist(workspace, uuid string) (*types.Checklist, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetChecklist(ws.Checklists, uuid)
}

func (J *JSONFile) GetAttachment(workspace, uuid string) (*types.Attachment, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetAttachment(ws.Attachments, uuid)
}

func (J *JSONFile) GetDate(workspace, uuid string) (*types.Date, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetDate(ws.Dates, uuid)
}
