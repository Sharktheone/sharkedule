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
	Text          Type = "text"
	MultilineText Type = "multilineText"
	Number        Type = "number"
	Date          Type = "date"
	MultiList     Type = "multiList"
	SingleList    Type = "singleList"
	Iteration     Type = "iteration"
	Checkbox      Type = "checkbox"
	Link          Type = "link"
	Email         Type = "email"
	Phone         Type = "phone"
	Location      Type = "location"
	Progress      Type = "progress"
	File          Type = "file"
	Dependency    Type = "dependency"
	Dependents    Type = "dependents"
	Assignee      Type = "assignee"
	Assignees     Type = "assignees"
	Rating        Type = "rating"
	Time          Type = "time"
	Team          Type = "team"
	Color         Type = "color"
	Month         Type = "month"
	Week          Type = "week"
	Year          Type = "year"
	Day           Type = "day"
	Range         Type = "range"

	Action Type = "action"

	//Element types
	Element Type = "element"
	Project Type = "project"
	Group   Type = "group"
	Board   Type = "board"
	Column  Type = "column"
	Task    Type = "task"
	Event   Type = "event"
	//Workspace     Type = "workspace" - not sure about this one

)
