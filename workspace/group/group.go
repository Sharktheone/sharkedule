package group

import (
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
)

type Group struct {
	UUID        string   `json:"uuid" bson:"uuid" yaml:"uuid"`
	Name        string   `json:"name" bson:"name" yaml:"name"`
	Description string   `json:"description" bson:"description" yaml:"description"`
	Icon        string   `json:"icon" bson:"icon" yaml:"icon"`
	IconColor   string   `json:"icon_color" bson:"icon_color" yaml:"icon_color"`
	Items       []string `json:"items" bson:"items" yaml:"items"`

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

func (g *Group) GetDescription() string {
	return g.Description
}

func (g *Group) GetIcon() string {
	return g.Icon
}

func (g *Group) GetIconColor() string {
	return g.IconColor
}

func (g *Group) GetItems() []string {
	return g.Items
}

func (g *Group) AddItem(item string) {
	g.Items = append(g.Items, item)
}

func (g *Group) RemoveItem(item string) {
	for i, it := range g.Items {
		if it == item {
			g.Items = append(g.Items[:i], g.Items[i+1:]...)
			return
		}
	}
}

func (g *Group) UpdateIcon(icon string) {
	g.Icon = icon
}

func (g *Group) UpdateIconColor(color string) {
	g.IconColor = color
}
