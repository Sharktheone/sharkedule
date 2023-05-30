package column

import (
	"sharkedule/kanban/KTypes/description"
	"sharkedule/kanban/board/column/task"
	"sync"
)

type Column struct {
	Mu          sync.Mutex               `json:"-" yaml:"-" bson:"-"`
	UUID        string                   `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                   `json:"name" yaml:"name" bson:"name"`
	Description *description.Description `json:"description,omitempty" yaml:"description" bson:"description"`
	Tasks       []*task.Task             `json:"tasks,omitempty" yaml:"tasks" bson:"tasks"`
	Index       int                      `json:"index" yaml:"index" bson:"index"`
	Board       string                   `json:"board" yaml:"board" bson:"board"`
}
