package KTypes

import "sync"

type CheckList struct {
	Mu    sync.Mutex      `json:"-" yaml:"-" bson:"-"`
	UUID  string          `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string          `json:"name" yaml:"name" bson:"name"`
	Items []CheckListItem `json:"items,omitempty" yaml:"items" bson:"items"`
}

type CheckListItem struct {
	Mu      sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID    string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name    string     `json:"name" yaml:"name" bson:"name"`
	Checked bool       `json:"checked" yaml:"checked" bson:"checked"`
}
