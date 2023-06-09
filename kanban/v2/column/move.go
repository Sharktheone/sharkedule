package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

func (c *Column) Move(board string, toIndex int) error {
	return db.DBV2.MoveColumn(board, c.UUID, toIndex)
}
