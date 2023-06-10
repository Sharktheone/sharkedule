package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

func (c *Column) Move(board string, toIndex int) error {
	return db.DB.MoveColumn(board, c.UUID, toIndex)
}
