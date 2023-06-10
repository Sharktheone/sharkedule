package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) Rename(name string) error {
	return db.DB.RenameTask(t.UUID, name)
}
