package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func GetTags(tags []*types.Tag) []*types.Tag {
	return tags
}

func GetTag(tags []*types.Tag, uuid string) (*types.Tag, error) {
	for _, t := range tags {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("tag with uuid %s does not exist", uuid)
}

func SaveTag(tags []*types.Tag, tag *types.Tag) error {
	for i, t := range tags {
		if t.UUID == tag.UUID {
			tags[i] = tag
			return nil
		}
	}
	return fmt.Errorf("tag with uuid %s does not exist", tag.UUID)
}

func SaveTags(tags []*types.Tag, tagsToSave []*types.Tag) {
	tags = tagsToSave
}

func RenameTag(tags []*types.Tag, uuid, name string) error {
	for _, t := range tags {
		if t.UUID == uuid {
			t.Name = name
			return nil
		}
	}
	return fmt.Errorf("tag with uuid %s does not exist", uuid)
}

func DeleteTag(tags []*types.Tag, tag string) error {
	for index, t := range tags {
		if t.UUID == tag {
			tags = append(tags[:index], tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting tag %s not found", tag)
}
