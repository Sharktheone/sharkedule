package workspace

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/google/uuid"
)

type Workspace struct {
	*types.Workspace

	//All uuids here as list?
}

type List struct {
	UUID        string
	Name        string
	Description string
	Cover       string
	Archived    bool
	Color       string
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Workspace: &types.Workspace{
			Name: name,
			UUID: uuid.New().String(),
		},
	}
}
