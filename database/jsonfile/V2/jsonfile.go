package jsonfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sharktheone/sharkedule/database"
	"github.com/Sharktheone/sharkedule/kanban/database"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
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
