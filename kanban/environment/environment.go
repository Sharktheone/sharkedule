package environment

import (
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

func GetBoardEnv(uuid *string) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{},
		boardUUIDs:  []*string{uuid},
	}
	env.Index()

	return env.Environment
}

func GetColumnEnv(uuid *string) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{},
		columnUUIDs: []*string{uuid},
	}
	env.Index()

	return env.Environment
}

func GetTaskEnv(uuid *string) *types2.Environment {
	env := Environment{
		Environment: &types2.Environment{},
		taskUUIDs:   []*string{uuid},
	}
	env.Index()

	return env.Environment
}
