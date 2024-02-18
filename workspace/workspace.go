package workspace

import (
	ktypes "github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/types"
	"github.com/google/uuid"
)

type Workspace struct {
	UUID   string        `json:"uuid"`
	Fields []types.Field `json:"fields"`
	Users  []string      `json:"users"`
}

type List struct {
	UUID        string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string `json:"name" yaml:"name" bson:"name"`
	Description string `json:"description" yaml:"description" bson:"description"`
	Cover       string `json:"cover" yaml:"cover" bson:"cover"`
	Archived    bool   `json:"archived" yaml:"archived" bson:"archived"`
	Color       string `json:"color" yaml:"color" bson:"color"`
}

type Info struct {
	*List
	Boards []*ktypes.NameList `json:"boards" yaml:"boards" bson:"boards"`
}

func NewWorkspace() *Workspace {
	return &Workspace{
		UUID: uuid.New().String(),
	}
}

func (w *Workspace) GetUUID() string {
	return w.UUID
}
