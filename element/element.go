package element

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/Sharktheone/sharkedule/kanban/activity"
	"github.com/Sharktheone/sharkedule/user"
)

type Element struct {
	UUID     string            `json:"uuid" bson:"uuid" yaml:"uuid"`
	Type     elements.Type     `json:"type" bson:"type" yaml:"type"`
	Fields   []field.Field     `json:"fields" bson:"fields" yaml:"fields"`
	Activity activity.Activity `json:"activity" bson:"activity" yaml:"activity"`
	user     *user.User        //used for access control, not serialized + may be nil
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
	return e.user.Access.DeleteElement(e.UUID)
}
