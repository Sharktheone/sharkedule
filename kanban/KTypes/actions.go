package KTypes

import "sync"

type Actions struct {
	Mu     sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID   string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name   string     `json:"name" yaml:"name" bson:"name"`
	Icon   string     `json:"icon" yaml:"icon" bson:"icon"`
	Color  string     `json:"color" yaml:"color" bson:"color"`
	Type   string     `json:"type" yaml:"type" bson:"type"`
	Action string     `json:"action" yaml:"action" bson:"action"`
}
