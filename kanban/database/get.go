package kanbandb

import (
	"fmt"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

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
