package database

import (
	"github.com/Sharktheone/sharkedule/types"
	"sync"
)

const (
	DBRoot = "./"
)

type DBStructure struct {
	Mu         *sync.Mutex       `json:"-" yaml:"-" bson:"-"`
	Workspaces []types.Workspace `json:"workspaces" yaml:"workspaces" bson:"workspaces"`
	Users      []types.User      `json:"users" yaml:"users" bson:"users"`
}

type IDatabase interface {
	// Load and Save
	Load() error
	Save() error

	//User functions
	GetUser(uuid string) (types.User, error)
	GetUserByMail(mail string) (types.User, error)
	UpdateUserUsername(uuid string, username string) error
	UpdateUserEmail(uuid string, email string) error
	UpdateUserPassword(uuid string, password string) error
	AddUserWorkspaceAccess(uuid, workspace string) error
	RemoveUserWorkspaceAccess(uuid, workspace string) error
	UpdateUserSettings(uuid string, settings types.Settings) error
	DeleteUser(uuid string) error

	//Workspace functions
	GetWorkspace(uuid string) (types.Workspace, error)
	DeleteWorkspace(uuid string) error

	//Element functions
	GetElement(workspace, uuid string) (types.Element, error)
	CreateElement(workspace string, e *types.ElementType, name string) (types.Element, error)
	GetElements(workspace string, uuids []string) ([]types.Element, error)
}
