package user

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/user/email"
	"github.com/Sharktheone/sharkedule/user/mfa"
	"github.com/Sharktheone/sharkedule/user/oauth"
	"github.com/Sharktheone/sharkedule/user/settings"
	"github.com/Sharktheone/sharkedule/utils"
	"slices"
)

type User struct {
	UUID         string
	Username     string
	Password     string
	OAuth        oauth.OAuth
	MFA          mfa.MFA
	Email        email.EMail
	Boards       []string
	Columns      []string
	CustomColors []string
	Settings     settings.Settings
}

func (u *User) SaveBoard(workspace string, board *types.Board) error {
	if !slices.Contains(u.Boards, board.UUID) {
		return fmt.Errorf("board with uuid %s does not exist", board.UUID)
	}
	return db.DB.SaveBoard(workspace, board)
}

func (u *User) SaveBoards(workspace string, boards []*types.Board) error {
	for _, board := range boards {
		if !slices.Contains(u.Boards, board.UUID) {
			return fmt.Errorf("board with uuid %s does not exist", board.UUID)
		}
	}
	return db.DB.SaveBoards(workspace, boards)
}

func (u *User) SaveColumn(workspace string, column *types.Column) error {
	if !utils.SliceHaveCommon(u.Boards, column.Boards) || !slices.Contains(u.Columns, column.UUID) {
		return fmt.Errorf("column with uuid %s does not exist", column.UUID)
	}
	return db.DB.SaveColumn(workspace, column)
}

func (u *User) SaveColumns(workspace string, columns []*types.Column) error {

	for _, column := range columns {
		if !utils.SliceHaveCommon(u.Boards, column.Boards) || !slices.Contains(u.Columns, column.UUID) {
			return fmt.Errorf("column with uuid %s does not exist", column.UUID)
		}
	}
	return db.DB.SaveColumns(workspace, columns)
}

func (u *User) SaveTask(workspace string, task *types.Task) error {
	if !utils.SliceHaveCommon(u.Boards, task.Boards) || !utils.SliceHaveCommon(u.Columns, task.Columns) {
		return fmt.Errorf("task with uuid %s does not exist", task.UUID)
	}
	return db.DB.SaveTask(workspace, task)
}

func (u *User) SaveTasks(workspace string, tasks []*types.Task) error {
	for _, task := range tasks {
		if !utils.SliceHaveCommon(u.Boards, task.Boards) || !utils.SliceHaveCommon(u.Columns, task.Columns) {
			return fmt.Errorf("task with uuid %s does not exist", task.UUID)
		}
	}
	return db.DB.SaveTasks(workspace, tasks)
}

func (u *User) CreateBoard(workspace string, name string) (*types.Board, error) {
	board := types.NewBoard(name)
	u.Boards = append(u.Boards, board.UUID)
	if err := db.DB.SaveBoard(workspace, board); err != nil {
		return nil, err
	}
	return board, nil
}

func (u *User) GetAllBoards(workspace string) ([]*types.Board, error) {
	return db.DB.GetBoards(workspace, u.Boards)
}

func (u *User) GetBoard(workspace, uuid string) (*types.Board, error) {
	if !slices.Contains(u.Boards, uuid) {
		return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
	}
	return db.DB.GetBoard(workspace, uuid)
}

func (u *User) GetBoards(workspace string, uuids []string) ([]*types.Board, error) {
	for _, uuid := range uuids {
		if !slices.Contains(u.Boards, uuid) {
			return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
		}
	}
	return db.DB.GetBoards(workspace, uuids)
}

func (u *User) GetAllBoardNames(workspace string) ([]*namelist.NameList, error) {
	return db.DB.GetBoardNames(workspace, u.Boards)
}

func (u *User) GetBoardNames(workspace string, uuids []string) ([]*namelist.NameList, error) {
	for _, uuid := range uuids {
		if !slices.Contains(u.Boards, uuid) {
			return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
		}
	}

	return db.DB.GetBoardNames(workspace, uuids)
}

func (u *User) GetColumn(workspace, uuid string) (*types.Column, error) {
	if !slices.Contains(u.Columns, uuid) {
		return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
	}
	return db.DB.GetColumn(workspace, uuid)
}

//func (u *User) GetColumns(uuids []string) ([]*types.Column, error) TODO

func (u *User) GetTask(workspace, uuid string) (*types.Task, error) {
	return db.DB.GetTask(workspace, uuid)
}

func (u *User) GetAllTags(workspace string) ([]*types.Tag, error) {
	return db.DB.GetAllTags(workspace) //TODO: Filter tags for users
}

func (u *User) GetTag(workspace, uuid string) (*types.Tag, error) {
	return db.DB.GetTag(workspace, uuid) //TODO: Filter tags for users
}

//func (u *User) GetTags(uuids []string) TODO

func (u *User) GetStatus(workspace, uuid string) (*types.Status, error) {
	return db.DB.GetStatus(workspace, uuid) //TODO: Filter status for users
}

func (u *User) GetPriority(workspace, uuid string) (*types.Priority, error) {
	return db.DB.GetPriority(workspace, uuid) //TODO: Filter priority for users
}

func (u *User) GetMember(uuid string) (*types.Member, error) {
	return db.DB.GetUser(uuid) //TODO: Filter member for users
}

func (u *User) GetChecklist(workspace, uuid string) (*types.Checklist, error) {
	return db.DB.GetChecklist(workspace, uuid) //TODO: Filter checklist for users
}

func (u *User) GetAttachment(workspace, uuid string) (*types.Attachment, error) {
	return db.DB.GetAttachment(workspace, uuid) //TODO: Filter attachment for users
}

func (u *User) GetDate(workspace, uuid string) (*types.Date, error) {
	return db.DB.GetDate(workspace, uuid) //TODO: Filter date for users
}

func (u *User) DeleteBoard(workspace, uuid string) error {
	if !slices.Contains(u.Boards, uuid) {
		return fmt.Errorf("board with uuid %s does not exist", uuid)
	}
	return db.DB.DeleteBoard(workspace, uuid)
}

func (u *User) DeleteColumn(workspace, uuid string) error {
	if !slices.Contains(u.Columns, uuid) {
		return fmt.Errorf("column with uuid %s does not exist", uuid)
	}
	return db.DB.DeleteColumn(workspace, uuid)
}

func (u *User) MoveColumn(workspace, board, uuid string, toIndex int) error {
	if !slices.Contains(u.Boards, board) {
		return fmt.Errorf("board with uuid %s does not exist", board)
	}
	if !slices.Contains(u.Columns, uuid) {
		return fmt.Errorf("column with uuid %s does not exist", uuid)
	}
	return db.DB.MoveColumn(workspace, board, uuid, toIndex)
}

func (u *User) DeleteTask(workspace, uuid string) error {
	return db.DB.DeleteTask(workspace, uuid) //TODO: Filter task for users
}

func (u *User) MoveTask(workspace, column, uuid, toColumn string, toIndex int) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	if !slices.Contains(u.Columns, toColumn) {
		return fmt.Errorf("column with uuid %s does not exist", toColumn)
	}
	return db.DB.MoveTask(workspace, column, uuid, toColumn, toIndex)
}

func (u *User) DeleteTaskOnColumn(workspace, column, uuid string) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.DeleteTaskOnColumn(workspace, column, uuid)
}

func (u *User) NewTask(workspace, column, name string) (*types.Task, error) {
	if !slices.Contains(u.Columns, column) {
		return nil, fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.NewTask(workspace, column, name)
}

func (u *User) NewColumn(workspace, board, name string) (*types.Column, error) {
	if !slices.Contains(u.Boards, board) {
		return nil, fmt.Errorf("board with uuid %s does not exist", board)
	}
	return db.DB.NewColumn(workspace, board, name)
}

func (u *User) AddTagToTask(workspace, task, tag string) error {
	return db.DB.AddTagToTask(workspace, task, tag) //TODO: Filter task for users
}

func (u *User) RemoveTagOnTask(workspace, column, uuid string) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.RemoveTagOnTask(workspace, column, uuid)
}

func (u *User) SetTagsOnTask(workspace, task string, tags []string) error {
	return db.DB.SetTagsOnTask(workspace, task, tags) //TODO: Filter task for users
}

func (u *User) SetTaskDescription(workspace, task, description string) error {
	return db.DB.SetTaskDescription(workspace, task, description) //TODO: Filter task for users
}

func (u *User) RenameTask(workspace, task, name string) error {
	return db.DB.RenameTask(workspace, task, name) //TODO: Filter task for users
}

func (u *User) DeleteColumnOnBoard(workspace, board, column string) error {
	if !slices.Contains(u.Boards, board) {
		return fmt.Errorf("board with uuid %s does not exist", board)
	}
	return db.DB.DeleteColumnOnBoard(workspace, board, column)
}

func (u *User) RenameColumn(workspace, column, name string) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.RenameColumn(workspace, column, name)
}
