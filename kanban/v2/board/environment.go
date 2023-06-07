package board

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/environment"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func (b *Board) Env() *types2.Environment {
	return environment.GetBoardEnv(b)
}
