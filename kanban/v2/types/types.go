package types

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/activity"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
	"github.com/Sharktheone/sharkedule/kanban/v2/comment"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
)

type Board struct {
	Name        string              `json:"name"`
	UUID        string              `json:"uuid"`
	Columns     []string            `json:"columns"`
	Tags        []string            `json:"tags"`
	Description string              `json:"description"`
	Members     []string            `json:"members"`
	Priority    string              `json:"priority"`
	Status      string              `json:"status"`
	DueDate     string              `json:"due_date"`
	Dates       []string            `json:"dates"`
	Comments    []comment.Comment   `json:"comments"`
	Attachments []string            `json:"attachments"`
	Checklists  []string            `json:"checklists"`
	Archived    bool                `json:"archived"`
	Activity    []activity.Activity `json:"activity"`
	Actions     []string            `json:"actions"`
}

type Tag struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Member struct {
	Username       string `json:"username"`
	UUID           string `json:"uuid"`
	ProfilePicture string `json:"profile_picture"`
}

type Checklist struct {
	Name        string          `json:"name"`
	UUID        string          `json:"uuid"`
	Items       []ChecklistItem `json:"items"`
	Description string          `json:"description"`
}

type ChecklistItem struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Checked bool   `json:"checked"`
}

type Priority struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

type Status struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

type Attachment struct {
	UUID        string `json:"uuid"`
	User        string `json:"from"`
	Size        int64  `json:"size"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type Date struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Date        int64  `json:"date"`
	Description string `json:"description"`
}

type Environment struct {
	Tags           []*Tag                         `json:"tags"`
	Status         []*Status                      `json:"status"`
	Priority       []*Priority                    `json:"priority"`
	Columns        []*column.Column               `json:"columns"`
	Boards         []*Board                       `json:"boards"`
	Tasks          []*task.Task                   `json:"tasks"`
	Members        []*Member                      `json:"members"`
	Checklists     []*Checklist                   `json:"checklists"`
	Attachments    []*Attachment                  `json:"attachments"`
	Dates          []*Date                        `json:"dates"`
	BoardNames     map[string]string              `json:"board_names,omitempty"`     // boardUUID -> name
	ColumnNames    map[string]map[string]string   `json:"column_names,omitempty"`    // columnUUID -> boardUUID -> name
	DependentTasks map[string]map[string][]string `json:"dependent_tasks,omitempty"` // taskUUID -> boardUUID -> columnUUID
}
