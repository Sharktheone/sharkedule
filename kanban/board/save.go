package board

import "github.com/Sharktheone/sharkedule/database/db"

func (b *Board) Save() error {
	return db.DB.SaveBoard(b.Workspace, b.Board)
}
