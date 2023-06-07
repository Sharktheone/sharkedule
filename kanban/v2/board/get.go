package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

func GetBoard(uuid string) (*Board, error) {
	b, err := db.DB.GetBoard(uuid)
	if err != nil {
		return nil, err
	}
	return ConvertBoard(b)
}

func GetBoards() ([]*Board, error) {
	b, err := db.DB.GetBoards()
	if err != nil {
		return nil, err
	}
	return ConvertBoards(b)
}
