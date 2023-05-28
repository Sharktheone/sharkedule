package tag

import "sync"

type Tag struct {
	Mu    sync.Mutex `json:"-" yaml:"-" bson:"-"`
	UUID  string     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string     `json:"name" yaml:"name" bson:"name"`
	Color string     `json:"color" yaml:"color" bson:"color"`
	Icon  string     `json:"icon,omitempty" yaml:"icon" bson:"icon"`
}
