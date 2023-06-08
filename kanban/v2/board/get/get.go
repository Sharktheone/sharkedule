package getboard

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
)

func Get(uuid string) (*board.Board, error) {
	b, err := db.DBV2.GetBoard(uuid)
	if err != nil {
		return nil, err
	}

	return &board.Board{Board: b}, nil
	//if err != nil {
	//	return nil, err
	//}
	//return environment.GetBoardEnv(board), nil
}

func GetBoards() ([]*board.Board, error) {
	boards, er := db.DBV2.GetBoards()
	if er != nil {
		return nil, er
	}

	var bds []*board.Board
	for _, b := range boards {
		bds = append(bds, &board.Board{Board: b})
	}
	return bds, nil
}

func Names() ([]*namelist.NameList, error) {
	return db.DBV2.GetBoardNames()
}
