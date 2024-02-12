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
	GetActivity() int //TODO
	GetReferencedBy() ReferenceGroup
	GetReferences() ReferenceGroup
	GetUser() string
	SetUser(u string)
	GetWorkspace() string
	SetWorkspace(ws string)
}

type Workspace interface {
	GetUUID() string
	GetName() string
	GetElement(u string) Element
}

type Reference interface {
	GetUUID() string
	GetField() []string
	GetLinked() bool
}

type ReferenceGroup interface {
	AddReference(r Reference)
}
