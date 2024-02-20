package database

import (
	ktypes "github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/types"
	"github.com/Sharktheone/sharkedule/user/settings"
	"sync"
)

const (
	DBRoot = "./"
)

type DBStructure struct {
	Mu         *sync.Mutex        `json:"-" yaml:"-" bson:"-"`
	Workspaces []*types.Workspace `json:"workspaces" yaml:"workspaces" bson:"workspaces"`
	Users      []*ktypes.User     `json:"users" yaml:"users" bson:"users"`
}

type IDatabase interface {
	// Load and Save
	Load() error
	Save() error

	//User functions
	GetUser(uuid string) (*ktypes.User, error)
	GetUserByMail(mail string) (*ktypes.User, error)
	UpdateUserUsername(uuid string, username string) error
	UpdateUserEmail(uuid string, email string) error
	UpdateUserPassword(uuid string, password string) error
	AddUserWorkspaceAccess(uuid, workspace string) error
	RemoveUserWorkspaceAccess(uuid, workspace string) error
	UpdateUserSettings(uuid string, settings settings.Settings) error
	DeleteUser(uuid string) error

	//Workspace functions
	GetWorkspace(uuid string) (*types.Workspace, error)
	DeleteWorkspace(uuid string) error

	//Board functions
	CreateBoard(workspace, name string) (*ktypes.Board, error)
	SaveBoard(workspace string, board *ktypes.Board) error
	SaveBoards(workspace string, boards []*ktypes.Board) error
	GetBoard(workspace, uuid string) (*ktypes.Board, error)
	GetAllBoards(workspace string) ([]*ktypes.Board, error)
	GetBoards(workspace string, uuids []string) ([]*ktypes.Board, error)
	GetAllBoardNames(workspace string) ([]*types.NameList, error)
	GetBoardNames(workspace string, uuids []string) (names []*types.NameList, err error)
	DeleteBoard(workspace, uuid string) error

	//Column functions
	SaveColumn(workspace string, column *ktypes.Column) error
	SaveColumns(workspace string, columns []*ktypes.Column) error
	GetColumn(workspace, uuid string) (*ktypes.Column, error)
	DeleteColumnOnBoard(workspace, board, column string) error
	RenameColumn(workspace, column, name string) error
	DeleteColumn(workspace, uuid string) error
	MoveColumn(workspace, board, uuid string, toIndex int) error
	NewColumn(workspace, board, name string) (*ktypes.Column, error)

	//Task functions
	SaveTask(workspace string, task *ktypes.Task) error
	SaveTasks(workspace string, tasks []*ktypes.Task) error
	GetTask(workspace, uuid string) (*ktypes.Task, error)
	DeleteTaskOnColumn(workspace, column, uuid string) error
	DeleteTask(workspace, uuid string) error
	MoveTask(workspace, column, uuid, toColumn string, toIndex int) error
	NewTask(workspace, column, name string) (*ktypes.Task, error)
	RenameTask(workspace, task, name string) error
	RemoveTagOnTask(workspace, column, uuid string) error
	SetTagsOnTask(workspace, task string, tags []string) error
	SetTaskDescription(workspace, task, description string) error

	//Tag functions
	GetAllTags(workspace string) ([]*ktypes.Tag, error)
	GetTag(workspace, uuid string) (*ktypes.Tag, error)
	AddTagToTask(workspace, task, tag string) error

	//Other functions
	GetStatus(workspace, uuid string) (*ktypes.Status, error)
	GetPriority(workspace, uuid string) (*ktypes.Priority, error)
	//GetUser(uuid string) (*types.Member, error) TODO
	GetChecklist(workspace, uuid string) (*ktypes.Checklist, error)
	GetAttachment(workspace, uuid string) (*ktypes.Attachment, error)
	GetDate(workspace, uuid string) (*ktypes.Date, error)

	//Element functions
	GetElement(workspace, uuid string) (*types.Element, error)
	CreateElement(workspace string, e *types.ElementType, name string) (*types.Element, error)
	GetElements(workspace string, uuids []string) ([]*types.Element, error)
}
