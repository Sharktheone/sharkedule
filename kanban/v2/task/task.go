package task

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

type Task struct {
	*types.Task
}

func Get(uuid string) (*Task, error) {
	t, err := db.DBV2.GetTask(uuid)
	if err != nil {
		return nil, err
	}

	return &Task{
		Task: t,
	}, nil
}
