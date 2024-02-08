package jsonfile

import (
	"github.com/Sharktheone/sharkedule/element"
	"github.com/Sharktheone/sharkedule/kanban/database"
)

func (J *JSONFile) GetElement(workspace string, elementUUID string) (*element.Element, error) {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	e, err := kanbandb.GetElement(ws.Elements, elementUUID)
	if err != nil {
		return nil, err
	}
	e.SetWorkspace(nil) //TODO: set Workspace, for that replace types.Workspace with workspace.Workspace (currently import cycle)
	return e, nil
}
