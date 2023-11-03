package task

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/environment"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func Get(workspace, uuid string) (*Task, error) {
	t, err := db.DB.GetTask(workspace, uuid)
	if err != nil {
		return nil, err
	}

	return &Task{
		Task:      t,
		Workspace: workspace,
	}, nil
}

func (t *Task) Env() *types.Environment {
	return environment.GetTaskEnv(&t.UUID)
}
