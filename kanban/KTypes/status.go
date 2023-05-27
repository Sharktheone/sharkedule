package KTypes

import "sync"

type Status struct {
	Mu   sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID string     `json:"uuid" yaml:"uuid" bson:"uuid"`
}
