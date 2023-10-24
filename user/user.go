package user

import (
	"errors"
	"fmt"
	"github.com/Sharktheone/sharkedule/database/db"
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

var (
	NotAllBoardsFound = errors.New("didn't found all boards")
)

func (u *User) SaveBoard(board *types.Board) error {
	if !slices.Contains(u.Boards, board.UUID) {
		return fmt.Errorf("board with uuid %s does not exist", board.UUID)
	}
	return db.DB.SaveBoard(board)
}

func (u *User) SaveBoards(boards []*types.Board) error {
	for _, board := range boards {
		if !slices.Contains(u.Boards, board.UUID) {
			return fmt.Errorf("board with uuid %s does not exist", board.UUID)
		}
	}
	return db.DB.SaveBoards(boards)
}

func (u *User) SaveColumn(column *types.Column) error {
	if !utils.SliceHaveCommon(u.Boards, column.Boards) || !slices.Contains(u.Columns, column.UUID) {
		return fmt.Errorf("column with uuid %s does not exist", column.UUID)
	}
	return db.DB.SaveColumn(column)
}

func (u *User) SaveColumns(columns []*types.Column) error {

	for _, column := range columns {
		if !utils.SliceHaveCommon(u.Boards, column.Boards) || !slices.Contains(u.Columns, column.UUID) {
			return fmt.Errorf("column with uuid %s does not exist", column.UUID)
		}
	}
	return db.DB.SaveColumns(columns)
}

func (u *User) SaveTask(task *types.Task) error {
	if !utils.SliceHaveCommon(u.Boards, task.Boards) || !utils.SliceHaveCommon(u.Columns, task.Columns) {
		return fmt.Errorf("task with uuid %s does not exist", task.UUID)
	}
	return db.DB.SaveTask(task)
}

func (u *User) SaveTasks(tasks []*types.Task) error {
	for _, task := range tasks {
		if !utils.SliceHaveCommon(u.Boards, task.Boards) || !utils.SliceHaveCommon(u.Columns, task.Columns) {
			return fmt.Errorf("task with uuid %s does not exist", task.UUID)
		}
	}
	return db.DB.SaveTasks(tasks)
}

func (u *User) CreateBoard(name string) (*types.Board, error) {
	board := types.NewBoard(name)
	u.Boards = append(u.Boards, board.UUID)
	if err := db.DB.SaveBoard(board); err != nil {
		return nil, err
	}
	return board, nil
}

func (u *User) GetAllBoards() ([]*types.Board, error) {
	return db.DB.GetBoards(u.Boards)
}

func (u *User) GetBoard(uuid string) (*types.Board, error) {
	if !slices.Contains(u.Boards, uuid) {
		return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
	}
	return db.DB.GetBoard(uuid)
}

func (u *User) GetBoards(uuids []string) ([]*types.Board, error) {
	for _, uuid := range uuids {
		if !slices.Contains(u.Boards, uuid) {
			return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
		}
	}
	return db.DB.GetBoards(uuids)
}

func (u *User) GetAllBoardNames() ([]*types.NameList, error) {
	return db.DB.GetBoardNames(u.Boards)
}

func (u *User) GetBoardNames(uuids []string) ([]*types.NameList, error) {
	for _, uuid := range uuids {
		if !slices.Contains(u.Boards, uuid) {
			return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
		}
	}

	return db.DB.GetBoardNames(uuids)
}

func (u *User) GetColumn(uuid string) (*types.Column, error) {
	if !slices.Contains(u.Columns, uuid) {
		return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
	}
	return db.DB.GetColumn(uuid)
}

//func (u *User) GetColumns(uuids []string) ([]*types.Column, error) TODO

func (u *User) GetTask(uuid string) (*types.Task, error) {
	return db.DB.GetTask(uuid)
}

func (u *User) GetAllTags() ([]*types.Tag, error) {
	return db.DB.GetAllTags() //TODO: Filter tags for users
}

func (u *User) GetTag(uuid string) (*types.Tag, error) {
	return db.DB.GetTag(uuid) //TODO: Filter tags for users
}

//func (u *User) GetTags(uuids []string) TODO

func (u *User) GetStatus(uuid string) (*types.Status, error) {
	return db.DB.GetStatus(uuid) //TODO: Filter status for users
}

func (u *User) GetPriority(uuid string) (*types.Priority, error) {
	return db.DB.GetPriority(uuid) //TODO: Filter priority for users
}

func (u *User) GetMember(uuid string) (*types.Member, error) {
	return db.DB.GetMember(uuid) //TODO: Filter member for users
}

func (u *User) GetChecklist(uuid string) (*types.Checklist, error) {
	return db.DB.GetChecklist(uuid) //TODO: Filter checklist for users
}

func (u *User) GetAttachment(uuid string) (*types.Attachment, error) {
	return db.DB.GetAttachment(uuid) //TODO: Filter attachment for users
}

func (u *User) GetDate(uuid string) (*types.Date, error) {
	return db.DB.GetDate(uuid) //TODO: Filter date for users
}

func (u *User) DeleteBoard(uuid string) error {
	if !slices.Contains(u.Boards, uuid) {
		return fmt.Errorf("board with uuid %s does not exist", uuid)
	}
	return db.DB.DeleteBoard(uuid)
}

func (u *User) DeleteColumn(uuid string) error {
	if !slices.Contains(u.Columns, uuid) {
		return fmt.Errorf("column with uuid %s does not exist", uuid)
	}
	return db.DB.DeleteColumn(uuid)
}

func (u *User) MoveColumn(board, uuid string, toIndex int) error {
	if !slices.Contains(u.Boards, board) {
		return fmt.Errorf("board with uuid %s does not exist", board)
	}
	if !slices.Contains(u.Columns, uuid) {
		return fmt.Errorf("column with uuid %s does not exist", uuid)
	}
	return db.DB.MoveColumn(board, uuid, toIndex)
}

func (u *User) DeleteTask(uuid string) error {
	return db.DB.DeleteTask(uuid) //TODO: Filter task for users
}

func (u *User) MoveTask(column, uuid, toColumn string, toIndex int) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	if !slices.Contains(u.Columns, toColumn) {
		return fmt.Errorf("column with uuid %s does not exist", toColumn)
	}
	return db.DB.MoveTask(column, uuid, toColumn, toIndex)
}

func (u *User) DeleteTaskOnColumn(column, uuid string) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.DeleteTaskOnColumn(column, uuid)
}

func (u *User) NewTask(column, name string) (*types.Task, error) {
	if !slices.Contains(u.Columns, column) {
		return nil, fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.NewTask(column, name)
}

func (u *User) NewColumn(board, name string) (*types.Column, error) {
	if !slices.Contains(u.Boards, board) {
		return nil, fmt.Errorf("board with uuid %s does not exist", board)
	}
	return db.DB.NewColumn(board, name)
}

func (u *User) AddTagToTask(task, tag string) error {
	return db.DB.AddTagToTask(task, tag) //TODO: Filter task for users
}

func (u *User) RemoveTagOnTask(column, uuid string) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.RemoveTagOnTask(column, uuid)
}

func (u *User) SetTagsOnTask(task string, tags []string) error {
	return db.DB.SetTagsOnTask(task, tags) //TODO: Filter task for users
}

func (u *User) SetTaskDescription(task, description string) error {
	return db.DB.SetTaskDescription(task, description) //TODO: Filter task for users
}

func (u *User) RenameTask(task, name string) error {
	return db.DB.RenameTask(task, name) //TODO: Filter task for users
}

func (u *User) DeleteColumnOnBoard(board, column string) error {
	if !slices.Contains(u.Boards, board) {
		return fmt.Errorf("board with uuid %s does not exist", board)
	}
	return db.DB.DeleteColumnOnBoard(board, column)
}

func (u *User) RenameColumn(column, name string) error {
	if !slices.Contains(u.Columns, column) {
		return fmt.Errorf("column with uuid %s does not exist", column)
	}
	return db.DB.RenameColumn(column, name)
}
