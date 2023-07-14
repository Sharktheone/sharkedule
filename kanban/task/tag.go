package task

import "github.com/Sharktheone/sharkedule/database/db"

func (t *Task) AddTag(uuid string) error {
	return db.DB.AddTagToTask(t.UUID, uuid)
}

func (t *Task) RemoveTag(uuid string) error {
	return db.DB.RemoveTagOnTask(t.UUID, uuid)
}

func (t *Task) SetTags(uuids []string) error {
	return db.DB.SetTagsOnTask(t.UUID, uuids)
}

func (t *Task) SetDescription(description string) error {
	return db.DB.SetTaskDescription(t.UUID, description)
}
