package environment

import (
	"github.com/Sharktheone/sharkedule/kanban/v2/board"
	"github.com/Sharktheone/sharkedule/kanban/v2/column"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func GetBoardEnv(b *board.Board) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{
			Boards: []*board.Board{b},
		},
	}
	env.Index()

	return env.Environment
}

func GetColumnEnv(c *column.Column) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{
			Columns: []*column.Column{c},
		},
	}
	env.Index()

	return env.Environment
}

func GetTaskEnv(t *task.Task) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{
			Tasks: []*task.Task{t},
		},
	}
	env.Index()

	return env.Environment
}
