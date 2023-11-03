package column

import "github.com/Sharktheone/sharkedule/database/db"

func (c *Column) Delete() error {
	return db.DB.DeleteColumn(c.Workspace, c.UUID)
}

func (c *Column) DeleteOnBoard(board string) error {
	return db.DB.DeleteColumnOnBoard(c.Workspace, board, c.UUID)
}
