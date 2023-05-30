package types

import (
	"sharkedule/kanban/KTypes/description"
	"sync"
)

type Column struct {
	Mu          sync.Mutex               `json:"-" yaml:"-" bson:"-"`
	UUID        string                   `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                   `json:"name" yaml:"name" bson:"name"`
	Description *description.Description `json:"description,omitempty" yaml:"description" bson:"description"`
	Tasks       []*Task                  `json:"tasks,omitempty" yaml:"tasks" bson:"tasks"`
	Index       int                      `json:"-" yaml:"-" bson:"-"`
	Board       string                   `json:"-" yaml:"-" bson:"-"`
}
