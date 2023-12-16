package environment

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func GetBoardEnv(workspace string, uuid *string) *types.Environment {
	env := Environment{
		Environment: &types.Environment{},
		boardUUIDs:  []*string{uuid},
		workspace:   workspace,
	}
	env.Index()

	return env.Environment
}

func GetColumnEnv(workspace string, uuid *string) *types.Environment {
	env := Environment{
		Environment: &types.Environment{},
		columnUUIDs: []*string{uuid},
		workspace:   workspace,
	}
	env.Index()

	return env.Environment
}

func GetTaskEnv(workspace string, uuid *string) *types.Environment {
	env := Environment{
		Environment: &types.Environment{},
		taskUUIDs:   []*string{uuid},
		workspace:   workspace,
	}
	env.Index()

	return env.Environment
}
