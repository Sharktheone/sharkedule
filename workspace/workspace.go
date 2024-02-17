package workspace

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/google/uuid"
)

type Workspace struct {
	*types.Workspace
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
	Boards []*types.NameList `json:"boards" yaml:"boards" bson:"boards"`
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Workspace: &types.Workspace{
			Name: name,
			UUID: uuid.New().String(),
		},
	}
}

func (w *Workspace) GetName() string {
	return w.Name
}

func (w *Workspace) GetUUID() string {
	return w.UUID
}
