package environment

import (
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

func GetBoardEnv(uuid *string) *types2.Environment {
	env := Environment{
		boardUUIDs: []*string{uuid},
	}
	env.Index()

	return env.Environment
}

func GetColumnEnv(uuid *string) *types2.Environment {
	env := Environment{
		columnUUIDs: []*string{uuid},
	}
	env.Index()

	return env.Environment
}

func GetTaskEnv(uuid *string) *types2.Environment {
	env := Environment{
		taskUUIDs: []*string{uuid},
	}
	env.Index()

	return env.Environment
}
