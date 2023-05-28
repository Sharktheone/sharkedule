package member

import "sync"

type Member struct {
	Mu   sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID string     `json:"uuid" yaml:"uuid" bson:"uuid"`
}
