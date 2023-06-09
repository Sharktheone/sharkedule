package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) Delete() error {
	return db.DBV2.DeleteTask(t.UUID)
}
