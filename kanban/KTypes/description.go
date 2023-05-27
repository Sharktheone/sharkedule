package KTypes

import "sync"

type Description struct {
	Mu          sync.Mutex `json:"-" yaml:"-" bson:"-"`
	Description string     `json:"description,omitempty" yaml:"description" bson:"description"`
}
