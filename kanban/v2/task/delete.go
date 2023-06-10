package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) Delete() error {
	return db.DB.DeleteTask(t.UUID)
}

func (t *Task) DeleteOnColumn(column string) error {
	return db.DB.DeleteTaskOnColumn(column, t.UUID)
}
