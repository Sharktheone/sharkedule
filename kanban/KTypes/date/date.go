package date

import "sync"

type Date struct {
	Mu   sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string     `json:"name" yaml:"name" bson:"name"`
	Date string     `json:"date" yaml:"date" bson:"date"`
}
