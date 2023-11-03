package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

func (c *Column) Move(board string, toIndex int) error {
	return db.DB.MoveColumn(c.Workspace, board, c.UUID, toIndex)
}
