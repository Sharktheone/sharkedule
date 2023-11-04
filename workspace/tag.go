package workspace

import "github.com/Sharktheone/sharkedule/kanban/tag"

// Tag functions
func (w *Workspace) GetAllTags(workspace string) ([]*tag.Tag, error) {

}

func (w *Workspace) GetTag(workspace, uuid string) (*tag.Tag, error) {

}

func (w *Workspace) AddTagToTask(workspace, task, tag string) error {

}
