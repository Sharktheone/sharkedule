package KTypes

import "sync"

type Priority struct {
	Mu       sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID     string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Priority string     `json:"priority" yaml:"priority" bson:"priority"`
}
