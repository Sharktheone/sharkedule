package permissions

type UserPerms struct {
}

type WorkspacePerms struct {
	CreateBoards bool
	UpdateBoards bool
	DeleteBoards bool

	CreateColumns bool
	UpdateColumns bool
	DeleteColumns bool
}

type OrgPerms struct {
}

type BoardPerms struct {
	Update bool
	Delete bool
}

type ColumnPerms struct {
}

type TaskPerms struct {
}
