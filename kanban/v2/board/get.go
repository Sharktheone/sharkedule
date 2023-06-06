package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/environment"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func Get(uuid string) (*types2.Environment, error) {
	board, err := db.DBV2.GetBoard(uuid)
	if err != nil {
		return nil, err
	}
	return environment.GetBoardEnv(board), nil
}
