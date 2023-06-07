package board

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

func (b *Board) Save() error {
	board, err := b.Convert()
	if err != nil {
		return err
	}
	return db.DB.SaveBoard(board)
}
