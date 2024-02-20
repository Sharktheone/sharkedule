package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/environment"
	"github.com/Sharktheone/sharkedule/kanban/types"
	types2 "github.com/Sharktheone/sharkedule/types"
)

func Get(workspace, uuid string) (*Board, error) {
	b, err := db.DB.GetBoard(workspace, uuid)
	if err != nil {
		return nil, err
	}

	return &Board{Board: b, Workspace: workspace}, nil
}

func GetBoardsAll(workspace string) ([]*Board, error) {
	boards, er := db.DB.GetAllBoards(workspace)
	if er != nil {
		return nil, er
	}

	var bds []*Board
	for _, b := range boards {
		bds = append(bds, &Board{Board: b, Workspace: workspace})
	}
	return bds, nil
}

func GetBoards(workspace string, uuids []string) ([]*Board, error) {
	boards, er := db.DB.GetBoards(workspace, uuids)
	if er != nil {
		return nil, er
	}

	var bds []*Board
	for _, b := range boards {
		bds = append(bds, &Board{Board: b, Workspace: workspace})
	}
	return bds, nil
}

func AllNames(workspace string) ([]*types2.NameList, error) {
	return db.DB.GetAllBoardNames(workspace)
}

func Names(workspace string, uuids []string) ([]*types2.NameList, error) {
	return db.DB.GetBoardNames(workspace, uuids)
}

func (b *Board) Env() *types.Environment {
	return environment.GetBoardEnv(b.Workspace, &b.UUID)
}
