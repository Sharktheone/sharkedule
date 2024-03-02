package environment

import (
	ktypes "github.com/Sharktheone/sharkedule/kanban/types"
)

func GetBoardEnv(workspace string, uuid *string) *ktypes.Environment {
	env := Environment{
		Environment: &ktypes.Environment{},
		boardUUIDs:  []*string{uuid},
		workspace:   workspace,
	}
	env.Index()

	return env.Environment
}

func GetColumnEnv(workspace string, uuid *string) *ktypes.Environment {
	env := Environment{
		Environment: &ktypes.Environment{},
		columnUUIDs: []*string{uuid},
		workspace:   workspace,
	}
	env.Index()

	return env.Environment
}

func GetTaskEnv(workspace string, uuid *string) *ktypes.Environment {
	env := Environment{
		Environment: &ktypes.Environment{},
		taskUUIDs:   []*string{uuid},
		workspace:   workspace,
	}
	env.Index()

	return env.Environment
}
