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
	CustomColors []string
	Settings     settings.Settings
}

var (
	NotAllBoardsFound = errors.New("didn't found all boards")
)

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
	var knownBoards []string
	var allKnown error

	for _, uuid := range uuids {
		if slices.Contains(u.Boards, uuid) {
			knownBoards = append(knownBoards, uuid)
		} else {
			allKnown = NotAllBoardsFound
		}
	}

	if boards, err := db.DB.GetBoards(knownBoards); err != nil {
		return boards, allKnown
	} else {
		return boards, err
	}
}

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
	if !utils.SliceHaveCommon(u.Boards, column.Boards) {
		return fmt.Errorf("board with uuid %s does not exist", column.BoardUUID)
	}
	return db.DB.SaveColumn(column)
}

func (u *User) DeleteBoard(uuid string) error {
	if !slices.Contains(u.Boards, uuid) {
		return fmt.Errorf("board with uuid %s does not exist", uuid)
	}
	return db.DB.DeleteBoard(uuid)
}
