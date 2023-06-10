package column

import "github.com/Sharktheone/sharkedule/database/db"

func (c *Column) Delete() error {
	return db.DB.DeleteColumn(c.UUID)
}

func (c *Column) DeleteOnBoard(board string) error {
	return db.DB.DeleteColumnOnBoard(board, c.UUID)
}
