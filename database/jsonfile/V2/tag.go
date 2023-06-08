package jsonfileV2

import (
	kanbandb "github.com/Sharktheone/sharkedule/kanban/v2/database"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func (J *JSONFile) GetTag(uuid string) (*types2.Tag, error) {
	return kanbandb.GetTag(J.db.Tags, uuid)
}

func (J *JSONFile) GetTags() ([]*types2.Tag, error) {
	return kanbandb.GetTags(J.db.Tags), nil
}

func (J *JSONFile) DeleteTag(uuid string) error {
	if err := kanbandb.DeleteTag(J.db.Tags, uuid); err != nil {
		return err
	}
	return J.Save()
}
