package getboard

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
)

func Get(uuid string) (*board.Board, error) {
	return db.DBV2.GetBoard(uuid)
	//if err != nil {
	//	return nil, err
	//}
	//return environment.GetBoardEnv(board), nil
}
