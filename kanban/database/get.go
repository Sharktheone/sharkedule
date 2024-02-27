package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/types"
)

func GetWorkspace(workspaces []*types.Workspace, uuid string) (*types.Workspace, error) {
	for _, w := range workspaces {
		if w.GetUUID() == uuid {
			return w, nil
		}
	}
	return nil, fmt.Errorf("workspace with uuid %s does not exist", uuid)
}

func DeleteWorkspace(workspaces []*types.Workspace, uuid string) error {
	for i, w := range workspaces {
		if w.GetUUID() == uuid {
			workspaces = append(workspaces[:i], workspaces[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("workspace with uuid %s does not exist", uuid)
}

func GetElement(elements []*types.Element, uuid string) (*types.Element, error) {
	for _, e := range elements {
		if e.GetUUID() == uuid {
			return e, nil
		}
	}
	return nil, fmt.Errorf("element with uuid %s does not exist", uuid)
}

func CreateElement(elements []*types.Element, elementType *types.ElementType, name string) (*types.Element, error) {
	var e *types.Element //TODO
	elements = append(elements, e)
	return e, nil
}

func GetElements(elements []*types.Element, uuids []string) ([]*types.Element, error) {
	var es []*types.Element
	for _, uuid := range uuids {
		e, err := GetElement(elements, uuid)
		if err != nil {
			return nil, err
		}
		es = append(es, e)
	}
	return es, nil
}
