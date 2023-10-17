package jsonfile

import (
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) GetTag(uuid string) (*types.Tag, error) {
	return kanbandb.GetTag(J.db.Tags, uuid)
}

func (J *JSONFile) GetTags() ([]*types.Tag, error) {
	return kanbandb.GetTags(J.db.Tags), nil
}

func (J *JSONFile) DeleteTag(uuid string) error {
	if err := kanbandb.DeleteTag(J.db.Tags, uuid); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) SaveTag(tag *types.Tag) error {
	if err := kanbandb.SaveTag(J.db.Tags, tag); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) SaveTags(tags []*types.Tag) error {
	kanbandb.SaveTags(J.db.Tags, tags)
	return J.Save()
}

func (J *JSONFile) RenameTag(uuid, name string) error {
	if err := kanbandb.RenameTag(J.db.Tags, uuid, name); err != nil {
		return err
	}
	return J.Save()
}
