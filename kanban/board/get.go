package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/environment"
	"github.com/Sharktheone/sharkedule/kanban/namelist"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func Get(uuid string) (*Board, error) {
	b, err := db.DB.GetBoard(uuid)
	if err != nil {
		return nil, err
	}

	return &Board{Board: b}, nil
}

func GetBoards() ([]*Board, error) {
	boards, er := db.DB.GetAllBoards()
	if er != nil {
		return nil, er
	}

	var bds []*Board
	for _, b := range boards {
		bds = append(bds, &Board{Board: b})
	}
	return bds, nil
}

func Names() ([]*namelist.NameList, error) {
	return db.DB.GetAllBoardNames()
}

func (b *Board) Env() *types.Environment {
	return environment.GetBoardEnv(&b.UUID)
}
