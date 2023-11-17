package permissions

type UserPerms struct {
	CreateWorkspaces bool
}

type WorkspacePerms struct {
	CreateBoards bool
	UpdateBoards bool
	DeleteBoards bool

	CreateColumns        bool
	UpdateColumns        bool
	DeleteColumns        bool
	RemoveColumnsOnBoard bool
	RenameColumns        bool
	MoveColumns          bool

	CreateTasks           bool
	UpdateTasks           bool
	DeleteTasks           bool
	MoveTasks             bool
	RenameTasks           bool
	RemoveTasksOnColumn   bool
	UpdateTaskDescription bool

	UpdateTagsOnTask bool
}

type OrgPerms struct {
}

type BoardPerms struct {
	Update        bool
	Delete        bool
	RemoveColumns bool
	MoveColumns   bool
	CreateColumns bool
}

type ColumnPerms struct {
	Update          bool
	Delete          bool
	RemoveFromBoard bool
	Rename          bool
	Move            bool
	CreateTasks     bool
	RemoveTasks     bool
	MoveTasks       bool
}

type TaskPerms struct {
	Update            bool
	Delete            bool
	RemoveFromColumn  bool
	Move              bool
	Rename            bool
	UpdateTags        bool
	UpdateDescription bool
}

type TagPerms struct {
	Update   bool
	Delete   bool
	UpdateOn bool
}

type StatusPerms struct {
}

type PriorityPerms struct {
}

type ChecklistPerms struct {
}

type AttachmentPerms struct {
}

type DatePerms struct {
}
