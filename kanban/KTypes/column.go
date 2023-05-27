package KTypes

import "sync"

type Column struct {
	Mu          sync.Mutex  `json:"-" yaml:"-" bson:"-"`
	UUID        string      `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string      `json:"name" yaml:"name" bson:"name"`
	Description Description `json:"description,omitempty" yaml:"description" bson:"description"`
	Tasks       []Task      `json:"tasks,omitempty" yaml:"tasks" bson:"tasks"`
}
