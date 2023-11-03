package board

import "github.com/Sharktheone/sharkedule/database/db"

func (b *Board) Delete() error {
	return db.DB.DeleteBoard(b.Workspace, b.UUID)
}
