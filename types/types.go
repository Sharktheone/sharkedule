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
	ListFields() []NameList
	GetField(f string) Field
	DeleteScopedField(f string) error
	CreateScopedField(name string) error
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
	ListFields() []NameList
	InfoField(uuid string) (FieldInfo, error)
	CreateField(f string) (string, error)
	DeleteField(f string) (string, error)
	ListLinkedFields() []NameList
	LinkField(f, from, to string) (string, error)
	UnlinkField(f, from, to string) (string, error)
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
	GetWorkspace(uuid string) (Workspace, error)
	SetUser(u string)
	SetWorkspace(ws string)
	SetField(f string)
	SetElement(e string)

	WorkspaceGetUUID() string
	WorkspaceGetName() string
	WorkspaceGetElement(u string) (Element, error)
	WorkspaceGetAllElements() ([]Element, error)
	WorkspaceCreateElement(t *ElementType, name string) (Element, error)
	WorkspaceCreateWithFields(t *ElementType, name string, fields []Field) (Element, error)
	WorkspaceCreateElementOn(u string, t *ElementType, name string) (Element, error)
	WorkspaceCreateWithFieldsOn(u string, t *ElementType, name string, fields []Field) (Element, error)
	WorkspaceGetElements(u []string) ([]Element, error)
	WorkspaceListFields() []NameList
	WorkspaceInfoField(uuid string) (FieldInfo, error)
	WorkspaceCreateField(f string) (string, error)
	WorkspaceDeleteField(f string) (string, error)
	WorkspaceListLinkedFields() []NameList
	WorkspaceLinkField(f, from, to string) (string, error)
	WorkspaceUnlinkField(f, from, to string) (string, error)

	ElementGetUUID() string
	ElementGetType() *ElementType
	ElementUpdateType(t *ElementType) error
	ElementGetFields() []Field
	ElementListFields() []NameList
	ElementGetField(f string) Field
	ElementDeleteScopedField(f string) error
	ElementCreateScopedField(name string) error
	ElementUpdateField(f string, v string) error
	ElementGetActivity() int //TODO
	ElementGetReferencedBy() ReferenceGroup
	ElementGetReferences() ReferenceGroup
	ElementAddReference(r Reference)
	ElementAddReferencedBy(r Reference)
	ElementGetUser() string
	ElementSetUser(u string)
	ElementGetWorkspace() Workspace
	ElementGetWorkspaceUUID() string
	ElementSetWorkspace(ws string) //only for db
	ElementAttach(u string) error
	ElementDetach(u string) error
	ElementGetAttachments() []string
	ElementDelete() error
	ElementMove(ref, to string, index int) error
	ElementCopy(to string, index int) error
	ElementMoveIndex(index int) error
	ElementGetIndex() int
	ElementMoveElement(e, to string, index int) error
	ElementCopyElement(e, to string, index int) error
	ElementGetSubElements() []Element
	ElementGetSubElementsUUID() []NameList
	ElementGetSubElementsType(t ElementType) []Element
	ElementGetSubElementsTypeUUID(t ElementType) []NameList
	ElementGetRecSubElements() []Element
	ElementGetRecSubElementsUUID() []NameList
	ElementGetRecSubElementsType(t ElementType) []Element
	ElementGetRecSubElementsTypeUUID(t ElementType) []NameList

	FieldGetName() string
	FieldGetUUID() string
	FieldGetValue() string
	FieldGetProperty(p string) string
	FieldGetParentType() ElementType
	FieldGetFieldType() FieldType
}

type SubType struct {
	UUID string `json:"uuid"`
	Type string `json:"type"`
}

type FieldInfo struct {
	Name string
	Type FieldType
	//...
}
