package permissions

type UserPerms struct {
	CreateWorkspaces bool `json:"create_workspaces" yaml:"create_workspaces" bson:"create_workspaces"`
}

type WorkspacePerms struct {
	CreateBoards bool `json:"create_boards" yaml:"create_boards" bson:"create_boards"`
	UpdateBoards bool `json:"update_boards" yaml:"update_boards" bson:"update_boards"`
	DeleteBoards bool `json:"delete_boards" yaml:"delete_boards" bson:"delete_boards"`

	CreateColumns        bool `json:"create_columns" yaml:"create_columns" bson:"create_columns"`
	UpdateColumns        bool `json:"update_columns" yaml:"update_columns" bson:"update_columns"`
	DeleteColumns        bool `json:"delete_columns" yaml:"delete_columns" bson:"delete_columns"`
	RemoveColumnsOnBoard bool `json:"remove_columns_on_board" yaml:"remove_columns_on_board" bson:"remove_columns_on_board"`
	RenameColumns        bool `json:"rename_columns" yaml:"rename_columns" bson:"rename_columns"`
	MoveColumns          bool `json:"move_columns" yaml:"move_columns" bson:"move_columns"`

	CreateTasks           bool `json:"create_tasks" yaml:"create_tasks" bson:"create_tasks"`
	UpdateTasks           bool `json:"update_tasks" yaml:"update_tasks" bson:"update_tasks"`
	DeleteTasks           bool `json:"delete_tasks" yaml:"delete_tasks" bson:"delete_tasks"`
	MoveTasks             bool `json:"move_tasks" yaml:"move_tasks" bson:"move_tasks"`
	RenameTasks           bool `json:"rename_tasks" yaml:"rename_tasks" bson:"rename_tasks"`
	RemoveTasksOnColumn   bool `json:"remove_tasks_on_column" yaml:"remove_tasks_on_column" bson:"remove_tasks_on_column"`
	UpdateTaskDescription bool `json:"update_task_description" yaml:"update_task_description" bson:"update_task_description"`

	UpdateTagsOnTask bool `json:"update_tags_on_task" yaml:"update_tags_on_task" bson:"update_tags_on_task"`
}

type OrgPerms struct {
}

type BoardPerms struct {
	Update        bool `json:"update" yaml:"update" bson:"update"`
	Delete        bool `json:"delete" yaml:"delete" bson:"delete"`
	RemoveColumns bool `json:"remove_columns" yaml:"remove_columns" bson:"remove_columns"`
	MoveColumns   bool `json:"move_columns" yaml:"move_columns" bson:"move_columns"`
	CreateColumns bool `json:"create_columns" yaml:"create_columns" bson:"create_columns"`
}

type ColumnPerms struct {
	Update          bool `json:"update" yaml:"update" bson:"update"`
	Delete          bool `json:"delete" yaml:"delete" bson:"delete"`
	RemoveFromBoard bool `json:"remove_from_board" yaml:"remove_from_board" bson:"remove_from_board"`
	Rename          bool `json:"rename" yaml:"rename" bson:"rename"`
	Move            bool `json:"move" yaml:"move" bson:"move"`
	CreateTasks     bool `json:"create_tasks" yaml:"create_tasks" bson:"create_tasks"`
	RemoveTasks     bool `json:"remove_tasks" yaml:"remove_tasks" bson:"remove_tasks"`
	MoveTasks       bool `json:"move_tasks" yaml:"move_tasks" bson:"move_tasks"`
}

type TaskPerms struct {
	Update            bool `json:"update" yaml:"update" bson:"update"`
	Delete            bool `json:"delete" yaml:"delete" bson:"delete"`
	RemoveFromColumn  bool `json:"remove_from_column" yaml:"remove_from_column" bson:"remove_from_column"`
	Move              bool `json:"move" yaml:"move" bson:"move"`
	Rename            bool `json:"rename" yaml:"rename" bson:"rename"`
	UpdateTags        bool `json:"update_tags" yaml:"update_tags" bson:"update_tags"`
	UpdateDescription bool `json:"update_description" yaml:"update_description" bson:"update_description"`
}

type TagPerms struct {
	Update   bool `json:"update" yaml:"update" bson:"update"`
	Delete   bool `json:"delete" yaml:"delete" bson:"delete"`
	UpdateOn bool `json:"update_on" yaml:"update_on" bson:"update_on"`
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
