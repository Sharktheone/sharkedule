package jsonfile

import (
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) GetWorkspace(uuid string) (*types.Workspace, error) {
	return kanbandb.GetWorkspace(J.db.Workspaces, uuid)
}
