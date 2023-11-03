package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) Save() error {
	return db.DB.SaveTask(t.Workspace, t.Task)
}
