package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) Move(toIndex int, column, toColumn, board string) error {
	return db.DBV2.MoveTask(t.UUID, toIndex, column, toColumn, board) // TODO: add handler
}
