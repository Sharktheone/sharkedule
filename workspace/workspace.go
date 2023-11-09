package workspace

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/google/uuid"
)

type Workspace struct {
	*types.Workspace

	//All uuids here as list?
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Workspace: &types.Workspace{
			Name: name,
			UUID: uuid.New().String(),
		},
	}
}

//TODO implement all other functions like on user
