package getboard

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
)

func Get(uuid string) (*board.Board, error) {
	return db.DBV2.GetBoard(uuid)
	//if err != nil {
	//	return nil, err
	//}
	//return environment.GetBoardEnv(board), nil
}

func GetBoards() ([]*board.Board, error) {
	return db.DBV2.GetBoards()
}

func Names() ([]*namelist.NameList, error) {
	return db.DBV2.GetBoardNames()
}
