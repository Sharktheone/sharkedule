package KTypes

import "sync"

type Activity struct {
	Mu      sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID    string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Message string     `json:"message" yaml:"message" bson:"message"`
	Date    string     `json:"date" yaml:"date" bson:"date"`
	User    string     `json:"user" yaml:"user" bson:"user"`
}
