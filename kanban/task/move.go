package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) Move(column, toColumn string, toIndex int) error {
	return db.DB.MoveTask(column, t.UUID, toColumn, toIndex)
}
