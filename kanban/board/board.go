package board

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
)

type Board struct {
	*types.Board
	Workspace string
}
