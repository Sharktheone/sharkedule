package access

import (
	"github.com/Sharktheone/sharkedule/user/permissions"
)

type Access struct {
	Workspaces  []WorkspaceAccess
	Permissions permissions.UserPerms
	//...
}
