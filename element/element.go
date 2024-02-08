package element

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/kanban/activity"
	"github.com/Sharktheone/sharkedule/user"
	"github.com/Sharktheone/sharkedule/workspace"
)

type Element struct {
	UUID         string               `json:"uuid" bson:"uuid" yaml:"uuid"`
	Type         elements.Type        `json:"type" bson:"type" yaml:"type"`
	Fields       []field.Field        `json:"fields" bson:"fields" yaml:"fields"`
	Activity     activity.Activity    `json:"activity" bson:"activity" yaml:"activity"`
	ReferencedBy ReferenceGroup       `json:"attachments" bson:"attachments" yaml:"attachments"` //UUIDs of elements that reference this element
	References   ReferenceGroup       `json:"references" bson:"references" yaml:"references"`    //UUIDs of elements that this element reference
	user         *user.User           //used for access control, not serialized + may be nil - inserted by user.Access
	workspace    *workspace.Workspace //may be nil - inserted from DB
}

// GetUser gets the user from which actions are performed (access control)
func (e *Element) GetUser() (*user.User, error) {
	if e.user == nil {
		return nil, fmt.Errorf("user not set")
	}
	return e.user, nil
}

// SetUser sets the user from which actions are performed (access control)
func (e *Element) SetUser(u *user.User) {
	e.user = u
}

func (e *Element) GetWorkspace() (*workspace.Workspace, error) {
	if e.workspace == nil {
		return nil, fmt.Errorf("workspace not set")
	}
	return e.workspace, nil
}

func (e *Element) SetWorkspace(ws *workspace.Workspace) {
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

func (e *Element) Delete() error {
	return e.user.Access.DeleteElement(e.workspace.UUID, e.UUID)
}
