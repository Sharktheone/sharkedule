package task

import (
	"github.com/Sharktheone/sharkedule/database/db"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
)

type Task struct {
	Name         string           `json:"name"`
	UUID         string           `json:"uuid"`
	Boards       []string         `json:"boards"`
	Columns      []string         `json:"columns"`
	Tags         []string         `json:"tags"`
	Dependencies []string         `json:"dependencies"`
	Dependents   []string         `json:"dependents"`
	Comments     []types2.Comment `json:"comments"`
	Description  string           `json:"description"`
	Members      []string         `json:"members"`
	Priority     string           `json:"priority"`
	Status       string           `json:"status"`
	DueDate      string           `json:"due_date"`
	Dates        []string         `json:"dates"`
	Attachments  []string         `json:"attachments"`
	CheckList    []string         `json:"check_list"`
	Done         bool             `json:"done"`
	Activity     []string         `json:"activity"`
}

func Get(uuid string) (*Task, error) {
	return db.DBV2.GetTask(uuid)
}
