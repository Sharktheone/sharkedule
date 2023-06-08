package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
)

func Get(uuid string) (*Board, error) {
	b, err := db.DBV2.GetBoard(uuid)
	if err != nil {
		return nil, err
	}

	return &Board{Board: b}, nil
}

func GetBoards() ([]*Board, error) {
	boards, er := db.DBV2.GetBoards()
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
	return db.DBV2.GetBoardNames()
}