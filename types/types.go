package types

type ElementType string
type FieldType string

type Field interface {
	GetName() string
	GetUUID() string
	GetValue() string            //Value = text in textfield, name in select, etc
	GetProperty(p string) string //Property = color in status, Items of checklist, etc
	GetParentType() ElementType  //For example if the field is limited to a certain type of element, maybe this should be an array?
	GetFieldType() FieldType
}

type Element interface {
	GetUUID() string
	GetType() *ElementType
	GetFields() []Field
	GetField(f string) Field
	UpdateField(f string, v string) error
	GetActivity() int //TODO
	GetReferencedBy() ReferenceGroup
	GetReferences() ReferenceGroup
	AddReference(r Reference)
	AddReferencedBy(r Reference)
	GetUser() string
	SetUser(u string)
	GetWorkspace() string
	SetWorkspace(ws string) //only for db
}

type Workspace interface {
	GetUUID() string
	GetName() string
	GetElement(u string) (Element, error)
	GetAllElements() ([]Element, error)
	CreateElement(t *ElementType, name string) (Element, error)
	GetElements(u []string) ([]Element, error)
}

type Reference interface {
	GetUUID() string
	GetField() []string
	GetLinked() bool
}

type ReferenceGroup interface {
	AddReference(r Reference)
}

type User interface {
	GetUUID() string
	GetUsername() string
	SetUsername(u string)
	GetEmail() string
	SetEmail(e string)
	GetWorkspaces() []string
	GetSettings() Settings
	GetPassword() string
	TokenIsValid(token string) bool
	GetAccess() Access
	SetSettings(s Settings)
	SetSetting(s string, v string)
	SetPassword(p string)
}

type Settings interface {
	GetUUID() string
	GetUser() string
	GetSetting(s string) string
	SetSetting(s string, v string)
}

type Access interface {
	RemoveWorkspaceAccess(uuid string) error
	AddWorkspaceAccess(uuid string) error
	GetWorkspaces() []Workspace
	GetField(workspace string, uuid string) (Field, error)
	GetElement(workspace string, uuid string) (Element, error)
}
