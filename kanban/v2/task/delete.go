package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) Delete() error {
	return db.DBV2.DeleteTask(t.UUID)
}

func (t *Task) DeleteOnColumn(column string) error {
	return db.DBV2.DeleteTaskOnColumn(column, t.UUID)
}
