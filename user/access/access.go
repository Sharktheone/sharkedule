package access

import "github.com/Sharktheone/sharkedule/user/permissions"

// Hmm, do I really want to do this like this?
type Access struct {
	Boards map[string]BoardAccess
	Orgs   map[string]OrgAccess
	Create CreateAccess
	//...
}

type BoardAccess struct {
	Columns     map[string]ColumnAccess
	Permissions permissions.Board
	//...
}

type ColumnAccess struct {
	Tasks map[string]TaskAccess
}

type TaskAccess struct {
	Permissions permissions.Task
}

type OrgAccess struct {
	Permissions permissions.Org
}

type CreateAccess struct {
	Boards bool
	Orgs   bool
}
