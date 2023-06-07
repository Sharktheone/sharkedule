package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/database/types"
)

func GetColumn(board, column string) (*types.Column, error) {
	return db.DB.GetColumn(board, column)
}
