package element

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/kanban/activity"
)

type Element struct {
	UUID         string            `json:"uuid" bson:"uuid" yaml:"uuid"`
	Type         elements.Type     `json:"type" bson:"type" yaml:"type"`
	Fields       []field.Field     `json:"fields" bson:"fields" yaml:"fields"`
	Activity     activity.Activity `json:"activity" bson:"activity" yaml:"activity"`
	ReferencedBy ReferenceGroup    `json:"attachments" bson:"attachments" yaml:"attachments"` //UUIDs of elements that reference this element
	References   ReferenceGroup    `json:"references" bson:"references" yaml:"references"`    //UUIDs of elements that this element reference
	user         string            //used for access control, not serialized + may be nil - inserted by user.Access
	workspace    string            //may be nil - inserted from DB
}

// GetUser gets the user from which actions are performed (access control)
func (e *Element) GetUser() (string, error) {
	if e.user == "" {
		return "", fmt.Errorf("user not set")
	}
	return e.user, nil
}

// SetUser sets the user from which actions are performed (access control)
func (e *Element) SetUser(u string) {
	e.user = u
}

func (e *Element) GetWorkspace() (string, error) {
	if e.workspace == "" {
		return "", fmt.Errorf("workspace not set")
	}
	return e.workspace, nil
}

func (e *Element) SetWorkspace(ws string) {
	e.workspace = ws
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
