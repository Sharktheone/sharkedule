package workspace

import (
	"github.com/Sharktheone/sharkedule/types"
	"github.com/google/uuid"
)

type Workspace struct {
	UUID   string        `json:"uuid"`
	Fields []types.Field `json:"fields"`
	Users  []string      `json:"users"`
}

func NewWorkspace() *Workspace {
	return &Workspace{
		UUID: uuid.New().String(),
	}
}

func (w *Workspace) GetUUID() string {
	return w.UUID
}
