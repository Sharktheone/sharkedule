package archived

import "sync"

type Archived struct {
	Mu       sync.Mutex `json:"-" yaml:"-" bson:"-"`
	Archived bool       `json:"archived" yaml:"archived" bson:"archived"`
	Date     string     `json:"date" yaml:"date" bson:"date"`
	User     string     `json:"user" yaml:"user" bson:"user"`
}
