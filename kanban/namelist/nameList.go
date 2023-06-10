package namelist

import "sync"

type NameList struct {
	Mu   sync.Mutex `json:"-" yaml:"-" bson:"-"`
	Name string     `json:"name" yaml:"name" bson:"name"`
	UUID string     `json:"uuid" yaml:"uuid" bson:"uuid"`
}
