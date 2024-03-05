package jsonfile

import (
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/types"
)

func (J *JSONFile) GetElement(workspace string, elementUUID string) (types.Element, error) {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	elements, err := ws.GetAllElements()
	if err != nil {
		return nil, err
	}

	e, err := kanbandb.GetElement(elements, elementUUID)
	if err != nil {
		return nil, err
	}
	e.SetWorkspace(workspace)
	return e, nil
}

func (J *JSONFile) CreateElement(workspace string, elementType *types.ElementType, name string) (types.Element, error) {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	elements, err := ws.GetAllElements()
	if err != nil {
		return nil, err
	}
	e, err := kanbandb.CreateElement(elements, elementType, name)
	if err != nil {
		return nil, err
	}
	e.SetWorkspace(workspace)
	return e, nil
}

func (J *JSONFile) GetElements(workspace string, elementUUIDs []string) ([]types.Element, error) {
	J.db.Mu.Lock()
	defer J.db.Mu.Unlock()
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	elements, err := ws.GetAllElements()
	if err != nil {
		return nil, err
	}

	e, err := kanbandb.GetElements(elements, elementUUIDs)
	if err != nil {
		return nil, err
	}

	return e, nil
}
