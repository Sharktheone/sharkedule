package board

import (
	"sharkedule/database/db"
)

func GetBoard(uuid string) (*Board, error) {
	return db.DB.GetBoard(uuid)
}

func GetBoards() ([]*Board, error) {
	return db.DB.GetBoards()
}
