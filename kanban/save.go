package kanban

import "sharkedule/database/db"

func (b *Board) Save() error {
	return db.DB.SaveBoard(b)
}
