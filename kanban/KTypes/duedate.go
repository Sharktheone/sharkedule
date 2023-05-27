package KTypes

import "sync"

type DateDue struct {
	Mu   sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Date string     `json:"date" yaml:"date" bson:"date"`
}
