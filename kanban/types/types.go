package types

import (
	"github.com/Sharktheone/sharkedule/element"
	"github.com/Sharktheone/sharkedule/kanban/activity"
	"github.com/Sharktheone/sharkedule/kanban/comment"
	"github.com/Sharktheone/sharkedule/project"
	"github.com/Sharktheone/sharkedule/user/access/workspaceaccess"
	"github.com/Sharktheone/sharkedule/user/mfa"
	"github.com/Sharktheone/sharkedule/user/oauth"
	"github.com/Sharktheone/sharkedule/user/permissions"
	"github.com/Sharktheone/sharkedule/user/settings"
)

type Board struct { //Replaced by Element
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

type Column struct { //Replace by Element
	Name        string   `json:"name"`
	UUID        string   `json:"uuid"`
	Boards      []string `json:"boards,omitempty"`
	Tasks       []string `json:"tasks,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Task struct { //Replace by Element
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

type Tag struct { //Replaced by Field
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type Environment struct {
	Tags           []*Tag                         `json:"tags,omitempty"`
	Columns        []*Column                      `json:"columns,omitempty"`
	Boards         []*Board                       `json:"boards,omitempty"`
	Tasks          []*Task                        `json:"tasks,omitempty"`
	BoardNames     map[string]string              `json:"board_names,omitempty"`     // boardUUID -> name
	ColumnNames    map[string]map[string]string   `json:"column_names,omitempty"`    // columnUUID -> boardUUID -> name
	DependentTasks map[string]map[string][]string `json:"dependent_tasks,omitempty"` // taskUUID -> boardUUID -> columnUUID
}

type Workspace struct { //replaced by Workspace
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Cover       string `json:"cover,omitempty"`
	Archived    bool   `json:"archived,omitempty"`
	Color       string `json:"color,omitempty"`

	Members []string `json:"members,omitempty"`

	Tags     []*Tag             `json:"tags,omitempty"`
	Tasks    []*Task            `json:"tasks,omitempty"`
	Columns  []*Column          `json:"columns,omitempty"`
	Boards   []*Board           `json:"boards,omitempty"`
	Projects []*project.Project `json:"projects,omitempty"`

	Elements []*element.Element `json:"elements,omitempty"`

	//All uuids here as list?
}

type User struct { //replaced by User
	UUID     string            `json:"uuid"`
	Username string            `json:"username"`
	Name     string            `json:"name,omitempty"`
	Email    string            `json:"email"`
	Password string            `json:"password"`
	OAuth    oauth.OAuth       `json:"oauth,omitempty"`
	MFA      mfa.MFA           `json:"mfa,omitempty"`
	Access   Access            `json:"access"`
	Settings settings.Settings `json:"settings"`
}

type Access struct { //replaced by Access
	Workspaces  []workspaceaccess.WorkspaceAccess `json:"workspaces"`
	Permissions permissions.UserPerms             `json:"permissions"`
	//...
}
