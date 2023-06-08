package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

func NewBoard(name string) (*Board, error) {
	b, err := db.DBV2.CreateBoard(name)
	if err != nil {
		return nil, err
	}

	return &Board{Board: b}, nil
}
