package types

type ElementType string
type FieldType string

type Field interface {
	GetName() string
	GetUUID() string
	GetValue() string
	GetParentType() ElementType
	GetFieldType() FieldType
}

type Element interface {
	GetUUID() string
	GetType() ElementType
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
	GetElement(u string) Element
	CreateElement(t ElementType, name string) (*Element, error)
	GetElements(u []string) []Element
}

type Reference interface {
	GetUUID() string
	GetField() []string
	GetLinked() bool
}

type ReferenceGroup interface {
	AddReference(r Reference)
}
