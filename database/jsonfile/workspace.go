package jsonfile

import (
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/workspace"
)

func (J *JSONFile) GetWorkspace(uuid string) (*workspace.Workspace, error) {
	return kanbandb.GetWorkspace(J.db.Workspaces, uuid)
}
