package access

import (
	"errors"
	"github.com/Sharktheone/sharkedule/user/permissions"
)

type Access struct {
	Workspaces  []WorkspaceAccess
	Permissions permissions.UserPerms
	//...
}

type WorkspaceAccess struct {
	UUID        string
	Permissions permissions.WorkspacePerms

	AllBoards bool
	Boards    []BoardAccess

	AllOrgs bool
	Orgs    []OrgAccess
}

type BoardAccess struct {
	UUID        string
	AllColumns  bool
	Columns     []ColumnAccess
	Permissions permissions.BoardPerms //TODO: Override permissions for specific boards
}

type ColumnAccess struct {
	UUID        string
	AllTasks    bool
	Tasks       []TaskAccess
	Permissions permissions.ColumnPerms //TODO: Override permissions for specific columns
}

type TaskAccess struct {
	UUID        string
	Permissions permissions.TaskPerms //TODO: Override permissions for specific tasks
}

type OrgAccess struct {
	UUID        string
	Permissions permissions.OrgPerms
}

type CreateAccess struct {
	Boards bool
	Orgs   bool
}

func (wa *WorkspaceAccess) board(uuid string) (*BoardAccess, error) {

	for _, b := range wa.Boards {
		if b.UUID == uuid {
			return &b, nil
		}
	}
	return nil, errors.New("board not found")
}
