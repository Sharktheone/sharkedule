package kanbandb

import (
	"fmt"
	ktypes "github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/types"
)

func GetStatus(status []*ktypes.Status, uuid string) (*ktypes.Status, error) {
	for _, s := range status {
		if s.UUID == uuid {
			return s, nil
		}
	}
	return nil, fmt.Errorf("status with uuid %s does not exist", uuid)
}

func GetPriority(priorities []*ktypes.Priority, uuid string) (*ktypes.Priority, error) {
	for _, p := range priorities {
		if p.UUID == uuid {
			return p, nil
		}
	}
	return nil, fmt.Errorf("priority with uuid %s does not exist", uuid)
}

//func GetUser(members []*types.Member, uuid string) (*types.Member, error) {
//	for _, m := range members {
//		if m.UUID == uuid {
//			return m, nil
//		}
//	}
//	return nil, fmt.Errorf("member with uuid %s does not exist", uuid)
//}

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
