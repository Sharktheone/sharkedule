package attachment

import "sync"

type Attachment struct {
	Mu   sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string     `json:"name" yaml:"name" bson:"name"`
	Type string     `json:"type" yaml:"type" bson:"type"`
	Size string     `json:"size" yaml:"size" bson:"size"`
	Date string     `json:"date" yaml:"date" bson:"date"`
	User string     `json:"user" yaml:"user" bson:"user"`
}
