package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func DeleteTag(tags []*types.Tag, tag string) error {
	for index, t := range tags {
		if t.UUID == tag {
			tags = append(tags[:index], tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting tag %s not found", tag)
}

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
