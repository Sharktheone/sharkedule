package board

import "sharkedule/database/types"

func (b *Board) Convert() *types.Board {
	return b.Board
}
