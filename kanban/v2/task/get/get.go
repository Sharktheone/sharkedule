package gettask

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/v2/task"
)

func Get(uuid string) (*task.Task, error) {
	return db.DBV2.GetTask(uuid)
}
