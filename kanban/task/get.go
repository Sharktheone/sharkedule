package task

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/environment"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func Get(uuid string) (*Task, error) {
	t, err := db.DB.GetTask(uuid)
	if err != nil {
		return nil, err
	}

	return &Task{
		Task: t,
	}, nil
}

func (t *Task) Env() *types.Environment {
	return environment.GetTaskEnv(&t.UUID)
}
