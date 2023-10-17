package environment

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func GetBoardEnv(uuid *string) *types.Environment {
	env := Environment{
		Environment: &types.Environment{},
		boardUUIDs:  []*string{uuid},
	}
	env.Index()

	return env.Environment
}

func GetColumnEnv(uuid *string) *types.Environment {
	env := Environment{
		Environment: &types.Environment{},
		columnUUIDs: []*string{uuid},
	}
	env.Index()

	return env.Environment
}

func GetTaskEnv(uuid *string) *types.Environment {
	env := Environment{
		Environment: &types.Environment{},
		taskUUIDs:   []*string{uuid},
	}
	env.Index()

	return env.Environment
}
