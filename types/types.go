package types

type ElementType string
type FieldType string

func ElementTypeFromString(s string) (ElementType, error) {
	//TODO: Check if the string is a valid ElementType
	return ElementType(s), nil
}

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
	UpdateType(t *ElementType) error
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
	GetWorkspace() Workspace
	GetWorkspaceUUID() string
	SetWorkspace(ws string) //only for db
	Attach(u string) error
	Detach(u string) error
	GetAttachments() []string
	Delete() error
	Move(ref, to string, index int) error
	Copy(to string, index int) error
	MoveIndex(index int) error
	GetIndex() int
	MoveElement(e, to string, index int) error
	CopyElement(e, to string, index int) error
	GetSubElements() []Element
	GetSubElementsUUID() []NameList
	GetSubElementsType(t ElementType) []Element
	GetSubElementsTypeUUID(t ElementType) []NameList
	GetRecSubElements() []Element
	GetRecSubElementsUUID() []NameList
	GetRecSubElementsType(t ElementType) []Element
	GetRecSubElementsTypeUUID(t ElementType) []NameList
}

type Workspace interface {
	GetUUID() string
	GetName() string
	GetElement(u string) (Element, error)
	GetAllElements() ([]Element, error)
	CreateElement(t *ElementType, name string) (Element, error)
	CreateWithFields(t *ElementType, name string, fields []Field) (Element, error)
	CreateElementOn(u string, t *ElementType, name string) (Element, error)
	CreateWithFieldsOn(u string, t *ElementType, name string, fields []Field) (Element, error)
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

type WorkspaceInfoFields struct {
	UUID   string
	Name   string
	Fields []Field
}

type WorkspaceInfo struct {
	UUID     string
	Name     string
	BGColor  string
	BColor   string
	Elements []Element
	Cover    string
}

type Access interface {
	RemoveWorkspaceAccess(uuid string) error
	AddWorkspaceAccess(uuid string) error
	GetWorkspaces() []Workspace
	GetField(workspace string, uuid string) (Field, error)
	GetElement(workspace string, uuid string) (Element, error)
	ListWorkspaces() ([]Workspace, error)
	ListWorkspacesWithFields(fields []string) ([]WorkspaceInfoFields, error)
	WorkspaceInfo() ([]*WorkspaceInfo, error) //TODO
	ListWithFields(uuid string, fields []string) (*WorkspaceInfoFields, error)
	DeleteWorkspace(uuid string) error
}

type SubType struct {
	UUID string `json:"uuid"`
	Type string `json:"type"`
}
