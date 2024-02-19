package jsonfile

import (
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/types"
)

func (J *JSONFile) GetWorkspace(uuid string) (*types.Workspace, error) {
	return kanbandb.GetWorkspace(J.db.Workspaces, uuid)
}

func (J *JSONFile) DeleteWorkspace(uuid string) error {
	return kanbandb.DeleteWorkspace(J.db.Workspaces, uuid)
}
