package board

import "github.com/Sharktheone/sharkedule/database/db"

func (b *Board) Delete() error {
	return db.DBV2.DeleteBoard(b.UUID)
}
