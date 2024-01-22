package element

import (
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/kanban/activity"
)

type Element struct {
	UUID     string            `json:"uuid" bson:"uuid" yaml:"uuid"`
	Type     elements.Type     `json:"type" bson:"type" yaml:"type"`
	Fields   []field.Field     `json:"fields" bson:"fields" yaml:"fields"`
	Activity activity.Activity `json:"activity" bson:"activity" yaml:"activity"`
}

func (e *Element) GetUUID() string {
	return e.UUID
}

func (e *Element) GetType() elements.Type {
	return e.Type
}

func (e *Element) GetFields() []field.Field {
	return e.Fields
}
