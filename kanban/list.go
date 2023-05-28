package kanban

import (
	"sharkedule/database/db"
	"sharkedule/kanban/KTypes/namelist"
)

func List() ([]*Board, error) {
	return db.DB.GetBoards()
}

func ListNames() ([]*namelist.NameList, error) {
	return db.DB.GetBoardNames()
}
