package task

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

type Task struct {
	*types.Task
}

func Get(uuid string) (*Task, error) {
	return db.DBV2.GetTask(uuid)
}
