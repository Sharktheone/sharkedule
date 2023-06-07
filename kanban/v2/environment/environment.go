package environment

import types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"

func GetBoardEnv(board *types2.Board) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{
			Boards: []*types2.Board{board},
		},
	}
	env.Index()

	return env.Environment
}

func GetColumnEnv(column *types2.Column) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{
			Columns: []*types2.Column{column},
		},
	}
	env.Index()

	return env.Environment
}

func GetTaskEnv(task *types2.Task) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{
			Tasks: []*types2.Task{task},
		},
	}
	env.Index()

	return env.Environment
}
