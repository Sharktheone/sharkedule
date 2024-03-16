package access

import (
	"errors"
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/types"
	"github.com/Sharktheone/sharkedule/user/access/workspaceaccess"
)

func (a *Access) GetWorkspace(uuid string) (types.Workspace, error) {
	for _, w := range a.Workspaces {
		if w.UUID == uuid {
			return db.DB.GetWorkspace(uuid)
		}
	}
	return nil, errors.New("workspace not found")
}
func (a *Access) WorkspaceInfo() ([]*types.WorkspaceInfo, error) {
	var info []*types.WorkspaceInfo

	for _, w := range a.Workspaces {
		var ws, err = db.DB.GetWorkspace(w.UUID) // we don't need to check for permissions, because we already have them => saves time
		if err != nil {
			return nil, err //TODO: this could be problematic, because when we haven't synced the database and so maybe not removed the workspace from the user but from the database
		}

		elements, err := ws.GetAllElements()
		if err != nil {
			return nil, err
		}
		info = append(info, &types.WorkspaceInfo{
			UUID:     ws.GetUUID(),
			Name:     ws.GetName(),
			BGColor:  "",
			BColor:   "",
			Elements: elements,
			Cover:    "",
		})
	}

	return info, nil

}

//GetUser(uuid string) (*types.Member, error) TODO

func (a *Access) workspace(uuid string) (*workspaceaccess.WorkspaceAccess, error) {
	for _, w := range a.Workspaces {
		if w.UUID == uuid {
			return &w, nil
		}
	}
	return nil, errors.New("workspace not found")
}

func (a *Access) DeleteWorkspace(uuid string) error {
	ws, err := a.workspace(uuid)
	if err != nil {
		return err
	}

	if !ws.Permissions.DeleteWorkspace {
		return fmt.Errorf("no permissions to delete workspace %s", uuid)
	}

	return db.DB.DeleteWorkspace(uuid)
}

func (a *Access) GetElement(workspace string, uuid string) (types.Element, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllElements {
		_, err := ws.Element(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	return db.DB.GetElement(workspace, uuid) //TODO
}

func (a *Access) GetField(workspace string, uuid string) (*field.Field, error) {
	ws, err := a.workspace(workspace)
	if err != nil {
		return nil, err
	}

	if !ws.AllFields {
		_, err := ws.Field(uuid) //when it is in the slice, the person has access to it
		if err != nil {
			return nil, err
		}
	}

	//return db.DB.GetField(workspace, uuid) //TODO
	return nil, nil
}

func (a *Access) DeleteElement(workspace, uuid string) error {

	//_, err := a.workspace(workspace)
	//if err != nil {
	//	return err
	//}
	//
	////TODO
	//
	//if !ws.Permissions.DeleteElements {
	//	return fmt.Errorf("no permissions to delete element in workspace %s", workspace)
	//}
	//
	//elem, err := ws.Element(uuid)
	//if err == nil {
	//	if !elem.Permissions.Delete {
	//		return fmt.Errorf("no permissions to delete element %s", uuid)
	//	}
	//}
	//
	//return db.DB.DeleteElement(workspace, uuid)
	return nil
}
