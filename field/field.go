package field

import (
	"github.com/Sharktheone/sharkedule/element/elements"
)

type Field struct {
	Name       string `json:"name" yaml:"name" bson:"name"`
	UUID       string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Value      string `json:"value" yaml:"value" bson:"value"`
	ParentType elements.Type
	FieldType  Type `json:"type" yaml:"type" bson:"type"`
}

type Type string

const (
	String     Type = "string"
	Number     Type = "number"
	Date       Type = "date"
	MultiList  Type = "multiList"
	SingleList Type = "singleList"
	Iteration  Type = "iteration"
)
