package jsonfile

import (
	"github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func (J *JSONFile) GetTag(workspace, uuid string) (*types.Tag, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetTag(ws.Tags, uuid)
}

func (J *JSONFile) GetAllTags(workspace string) ([]*types.Tag, error) {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return nil, err
	}

	return kanbandb.GetTags(ws.Tags), nil
}

func (J *JSONFile) DeleteTag(workspace, uuid string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.DeleteTag(ws.Tags, uuid); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) SaveTag(workspace string, tag *types.Tag) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.SaveTag(ws.Tags, tag); err != nil {
		return err
	}
	return J.Save()
}

func (J *JSONFile) SaveTags(workspace string, tags []*types.Tag) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	kanbandb.SaveTags(ws.Tags, tags)
	return J.Save()
}

func (J *JSONFile) RenameTag(workspace, uuid, name string) error {
	ws, err := J.GetWorkspace(workspace)
	if err != nil {
		return err
	}

	if err := kanbandb.RenameTag(ws.Tags, uuid, name); err != nil {
		return err
	}
	return J.Save()
}
