package workspace

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/tag"
)

// Tag functions
func (w *Workspace) GetAllTags() ([]*tag.Tag, error) {
	tags, err := db.DB.GetAllTags(w.UUID)
	if err != nil {
		return nil, err
	}

	var tgs []*tag.Tag
	for _, t := range tags {
		tgs = append(tgs, &tag.Tag{Tag: t, Workspace: w.UUID})
	}
	return tgs, nil
}

func (w *Workspace) GetTag(uuid string) (*tag.Tag, error) {
	t, err := db.DB.GetTag(w.UUID, uuid)
	if err != nil {
		return nil, err
	}

	return &tag.Tag{Tag: t, Workspace: w.UUID}, nil

}

func (w *Workspace) AddTagToTask(task, tag string) error {
	return db.DB.AddTagToTask(w.UUID, task, tag)

}
