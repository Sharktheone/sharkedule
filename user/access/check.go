package access

import (
	"errors"
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/element"
	"github.com/Sharktheone/sharkedule/field"
	types2 "github.com/Sharktheone/sharkedule/types"
	"github.com/Sharktheone/sharkedule/user/access/workspaceaccess"
	"github.com/Sharktheone/sharkedule/workspace"
)

func (a *Access) GetWorkspace(uuid string) (*workspace.Workspace, error) {
	for _, w := range a.Workspaces {
		if w.UUID == uuid {
			ws, error := db.DB.GetWorkspace(uuid)
			if error != nil {
				return nil, error
			}
			return &workspace.Workspace{Workspace: ws}, nil
		}
	}
	return nil, errors.New("workspace not found")
}

func (a *Access) ListWorkspaces() ([]*workspace.List, error) {
	var list []*workspace.List

	for _, w := range a.Workspaces {
		var ws, err = db.DB.GetWorkspace(w.UUID) // we don't need to check for permissions, because we already have them => saves time
		if err != nil {
			return nil, err //TODO: this could be problematic, because when we haven't synced the database and so maybe not removed the workspace from the user but from the database
		}

		list = append(list, &workspace.List{
			UUID:        ws.UUID,
			Name:        ws.Name,
			Description: ws.Description,
			Cover:       ws.Cover,
			Archived:    ws.Archived,
			Color:       ws.Color,
		})
	}

	return list, nil
}

func (a *Access) WorkspaceInfo() ([]*workspace.Info, error) {
	var info []*workspace.Info

	for _, w := range a.Workspaces {
		var ws, err = db.DB.GetWorkspace(w.UUID) // we don't need to check for permissions, because we already have them => saves time
		if err != nil {
			return nil, err //TODO: this could be problematic, because when we haven't synced the database and so maybe not removed the workspace from the user but from the database
		}

		var boards []*types2.NameList
		if w.AllBoards {
			var b, err = db.DB.GetAllBoardNames(w.UUID)
			if err != nil {
				return nil, err
			}
			boards = b
		} else {
			var brds []string
			for _, b := range w.Boards {
				brds = append(brds, b.UUID)
			}
			var b, err = db.DB.GetBoardNames(w.UUID, brds)
			if err != nil {
				return nil, err
			}
			boards = b

		}
		if err != nil {
			return nil, err
		}

		info = append(info, &workspace.Info{
			List: &workspace.List{
				UUID:        ws.UUID,
				Name:        ws.Name,
				Description: ws.Description,
				Cover:       ws.Cover,
				Archived:    ws.Archived,
				Color:       ws.Color,
			},
			Boards: boards,
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

//func (a *Access) board(workspace, uuid string) (*BoardAccess, error) {
//	ws, err := a.workspace(workspace)
//	if err != nil {
//		return nil, err
//	}
//
//	return ws.board(uuid)
//}

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

func (a *Access) GetElement(workspace string, uuid string) (*element.Element, error) {
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
