package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
)

func List() ([]*Board, error) {
	b, err := db.DB.GetBoards()
	if err != nil {
		return nil, err
	}
	return ConvertBoards(b)
}

func ListNames() ([]*namelist.NameList, error) {
	return db.DB.GetBoardNames()
}
