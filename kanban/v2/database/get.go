package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/KTypes/namelist"
	types2 "github.com/Sharktheone/sharkedule/kanban/v2/types"
)

func GetBoard(boards []*types2.Board, uuid string) (*types2.Board, error) {
	for _, b := range boards {
		if b.UUID == uuid {
			return b, nil
		}
	}
	return nil, fmt.Errorf("board with uuid %s does not exist", uuid)
}

func GetBoards(boards []*types2.Board) []*types2.Board {
	return boards
}

func GetBoardNames(boards []*types2.Board) []*namelist.NameList {
	var names []*namelist.NameList
	for _, b := range boards {
		names = append(names, &namelist.NameList{
			Name: b.Name,
			UUID: b.UUID,
		})
	}
	return names
}

func GetColumn(columns []*types2.Column, uuid string) (*types2.Column, error) {
	for _, c := range columns {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("column with uuid %s does not exist", uuid)
}

func GetTask(tasks []*types2.Task, uuid string) (*types2.Task, error) {
	for _, t := range tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with uuid %s does not exist", uuid)
}

func GetTags(tags []*types2.Tag) []*types2.Tag {
	return tags
}

func GetTag(tags []*types2.Tag, uuid string) (*types2.Tag, error) {
	for _, t := range tags {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("tag with uuid %s does not exist", uuid)
}

func GetStatus(status []*types2.Status, uuid string) (*types2.Status, error) {
	for _, s := range status {
		if s.UUID == uuid {
			return s, nil
		}
	}
	return nil, fmt.Errorf("status with uuid %s does not exist", uuid)
}

func GetPriority(priorities []*types2.Priority, uuid string) (*types2.Priority, error) {
	for _, p := range priorities {
		if p.UUID == uuid {
			return p, nil
		}
	}
	return nil, fmt.Errorf("priority with uuid %s does not exist", uuid)
}

func GetMember(members []*types2.Member, uuid string) (*types2.Member, error) {
	for _, m := range members {
		if m.UUID == uuid {
			return m, nil
		}
	}
	return nil, fmt.Errorf("member with uuid %s does not exist", uuid)
}

func GetChecklist(checklists []*types2.Checklist, uuid string) (*types2.Checklist, error) {
	for _, c := range checklists {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("checklist with uuid %s does not exist", uuid)
}

func GetAttachment(attachments []*types2.Attachment, uuid string) (*types2.Attachment, error) {
	for _, a := range attachments {
		if a.UUID == uuid {
			return a, nil
		}
	}
	return nil, fmt.Errorf("attachment with uuid %s does not exist", uuid)
}

func GetDate(dates []*types2.Date, uuid string) (*types2.Date, error) {
	for _, date := range dates {
		if date.UUID == uuid {
			return date, nil
		}
	}
	return nil, fmt.Errorf("date with uuid %s does not exist", uuid)
}
