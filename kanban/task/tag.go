package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) AddTag(uuid string) error {

	return db.DB.AddTagToTask(t.UUID, uuid)
}

func (t *Task) RemoveTag(uuid string) error {

	return db.DB.RemoveTagOnTask(t.UUID, uuid)
}
