package types

import (
	"github.com/Sharktheone/sharkedule/kanban/activity"
	"github.com/Sharktheone/sharkedule/kanban/comment"
	"github.com/Sharktheone/sharkedule/project"
	"github.com/Sharktheone/sharkedule/user/permissions"
)

type Board struct {
	Name        string              `json:"name"`
	UUID        string              `json:"uuid"`
	Columns     []string            `json:"columns,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	Description string              `json:"description,omitempty"`
	Members     []string            `json:"members,omitempty"`
	Priority    string              `json:"priority,omitempty"`
	Status      string              `json:"status,omitempty"`
	DueDate     string              `json:"due_date,omitempty"`
	Dates       []string            `json:"dates,omitempty"`
	Comments    []comment.Comment   `json:"comments,omitempty"`
	Attachments []string            `json:"attachments,omitempty"`
	Checklists  []string            `json:"checklists,omitempty"`
	Archived    bool                `json:"archived,omitempty"`
	Activity    []activity.Activity `json:"activity,omitempty"`
	Actions     []string            `json:"actions,omitempty"`
}

type Column struct {
	Name        string   `json:"name"`
	UUID        string   `json:"uuid"`
	Boards      []string `json:"boards,omitempty"`
	Tasks       []string `json:"tasks,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Task struct {
	Name         string            `json:"name"`
	UUID         string            `json:"uuid"`
	Boards       []string          `json:"boards,omitempty"`
	Columns      []string          `json:"columns,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	Dependencies []string          `json:"dependencies,omitempty"`
	Dependents   []string          `json:"dependents,omitempty"`
	Comments     []comment.Comment `json:"comments,omitempty"`
	Description  string            `json:"description,omitempty"`
	Members      []string          `json:"members,omitempty"`
	Priority     string            `json:"priority,omitempty"`
	Status       string            `json:"status,omitempty"`
	DueDate      string            `json:"due_date,omitempty"`
	Dates        []string          `json:"dates,omitempty"`
	Attachments  []string          `json:"attachments,omitempty"`
	Checklists   []string          `json:"checklists,omitempty"`
	Done         bool              `json:"done,omitempty"`
	Activity     []string          `json:"activity,omitempty"`
}

type Tag struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type Member struct {
	Username       string `json:"username"`
	UUID           string `json:"uuid"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

type Checklist struct {
	Name        string          `json:"name"`
	UUID        string          `json:"uuid"`
	Items       []ChecklistItem `json:"items"`
	Description string          `json:"description,omitempty"`
}

type ChecklistItem struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Checked bool   `json:"checked"`
}

type Priority struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
}

type Status struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
}

type Attachment struct {
	UUID        string `json:"uuid"`
	User        string `json:"from"`
	Size        int64  `json:"size"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
}

type Date struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Date        int64  `json:"date"`
	Description string `json:"description,omitempty"`
}

type Stage struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
}

type Environment struct {
	Tags           []*Tag                         `json:"tags,omitempty"`
	Status         []*Status                      `json:"status,omitempty"`
	Priority       []*Priority                    `json:"priority,omitempty"`
	Columns        []*Column                      `json:"columns,omitempty"`
	Boards         []*Board                       `json:"boards,omitempty"`
	Tasks          []*Task                        `json:"tasks,omitempty"`
	Members        []*Member                      `json:"members,omitempty"`
	Checklists     []*Checklist                   `json:"checklists,omitempty"`
	Attachments    []*Attachment                  `json:"attachments,omitempty"`
	Dates          []*Date                        `json:"dates,omitempty"`
	BoardNames     map[string]string              `json:"board_names,omitempty"`     // boardUUID -> name
	ColumnNames    map[string]map[string]string   `json:"column_names,omitempty"`    // columnUUID -> boardUUID -> name
	DependentTasks map[string]map[string][]string `json:"dependent_tasks,omitempty"` // taskUUID -> boardUUID -> columnUUID
}

type Workspace struct {
	UUID        string
	Name        string
	Description string
	Cover       string
	Archived    bool
	Color       string

	Members []string

	Tags        []*Tag
	Tasks       []*Task
	Columns     []*Column
	Boards      []*Board
	Checklists  []*Checklist
	Priorities  []*Priority
	Statuses    []*Status
	Attachments []*Attachment
	Dates       []*Date
	Stages      []*Stage
	Projects    []*project.Project

	//All uuids here as list?
}

type User struct {
	UUID     string   `json:"uuid"`
	Username string   `json:"username"`
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	OAuth    OAuth    `json:"oauth,omitempty"`
	MFA      MFA      `json:"mfa,omitempty"`
	Access   Access   `json:"access"`
	Settings Settings `json:"settings"`
}

type OAuth struct {
}

type MFA struct {
}

type Access struct {
	Workspaces  []WorkspaceAccess `json:"workspaces"`
	Permissions UserPerms         `json:"permissions"`
}

type Settings struct {
}

type UserPerms struct {
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

type WorkspaceAccess struct {
	UUID        string
	Permissions permissions.WorkspacePerms

	//TODO: roles / groups

	AllBoards bool
	Boards    []BoardAccess

	AllColumns bool
	Columns    []ColumnAccess

	AllTasks bool
	Tasks    []TaskAccess

	AllTags bool
	Tags    []TagAccess

	AllOrgs bool
	Orgs    []OrgAccess

	AllStatuses bool
	Statuses    []StatusAccess

	AllPriorities bool
	Priorities    []PriorityAccess

	AllChecklists bool
	Checklists    []ChecklistAccess

	AllAttachments bool
	Attachments    []AttachmentAccess

	AllDates bool
	Dates    []DateAccess
}

type BoardAccess struct {
	UUID        string
	AllColumns  bool
	Columns     []ColumnAccess
	Permissions BoardPerms //TODO: Override permissions for specific boards
}

type ColumnAccess struct {
	UUID        string
	AllTasks    bool
	Tasks       []TaskAccess
	Permissions ColumnPerms //TODO: Override permissions for specific columns
}

type TaskAccess struct {
	UUID        string
	Permissions TaskPerms //TODO: Override permissions for specific tasks
}

type TagAccess struct {
	UUID        string
	Permissions TagPerms //TODO: Override permissions for specific tags
}

type OrgAccess struct {
	UUID        string
	Permissions OrgPerms
}

type CreateAccess struct {
	Boards bool
	Orgs   bool
}

type StatusAccess struct {
	UUID        string
	Permissions StatusPerms
}

type PriorityAccess struct {
	UUID        string
	Permissions PriorityPerms
}

type ChecklistAccess struct {
	UUID        string
	Permissions ChecklistPerms
}

type AttachmentAccess struct {
	UUID        string
	Permissions AttachmentPerms
}

type DateAccess struct {
	UUID        string
	Permissions DatePerms
}

//TODO: Maybe do this with interfaces?
