package group

import (
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
)

type Group struct {
	UUID        string `json:"uuid" bson:"uuid" yaml:"uuid"`
	Name        string `json:"name" bson:"name" yaml:"name"`
	Description string `json:"description" bson:"description" yaml:"description"`
	Icon        string `json:"icon" bson:"icon" yaml:"icon"`
	IconColor   string `json:"icon_color" bson:"icon_color" yaml:"icon_color"`

	Fields []field.Field `json:"fields" bson:"fields" yaml:"fields"`
}

func (g *Group) GetUUID() string {
	return g.UUID
}

func (g *Group) GetName() string {
	return g.Name
}

func (g *Group) GetType() elements.Type {
	return elements.Group
}

func (g *Group) GetFields() []field.Field {
	return g.Fields
}
