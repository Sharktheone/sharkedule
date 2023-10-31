package workspace

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/project"
	"github.com/google/uuid"
)

type Workspace struct {
	UUID        string
	Name        string
	Tags        []*types.Tag
	Tasks       []*types.Task
	Columns     []*types.Column
	Boards      []*types.Board
	Checklists  []*types.Checklist
	Priorities  []*types.Priority
	Statuses    []*types.Status
	Attachments []*types.Attachment
	Dates       []*types.Date
	Stages      []*types.Stage
	Projects    []*project.Project

	Members []string

	//All uuids here as list?
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Name: name,
		UUID: uuid.New().String(),
	}
}

//TODO implement all other functions like on user
