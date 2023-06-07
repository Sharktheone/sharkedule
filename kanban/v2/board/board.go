package board

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/activity"
	"github.com/Sharktheone/sharkedule/kanban/v2/comment"
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
