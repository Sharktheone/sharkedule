package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/workspace"
)

func GetStatus(status []*types.Status, uuid string) (*types.Status, error) {
	for _, s := range status {
		if s.UUID == uuid {
			return s, nil
		}
	}
	return nil, fmt.Errorf("status with uuid %s does not exist", uuid)
}

func GetPriority(priorities []*types.Priority, uuid string) (*types.Priority, error) {
	for _, p := range priorities {
		if p.UUID == uuid {
			return p, nil
		}
	}
	return nil, fmt.Errorf("priority with uuid %s does not exist", uuid)
}

func GetUser(members []*types.Member, uuid string) (*types.Member, error) {
	for _, m := range members {
		if m.UUID == uuid {
			return m, nil
		}
	}
	return nil, fmt.Errorf("member with uuid %s does not exist", uuid)
}

func GetChecklist(checklists []*types.Checklist, uuid string) (*types.Checklist, error) {
	for _, c := range checklists {
		if c.UUID == uuid {
			return c, nil
		}
	}
	return nil, fmt.Errorf("checklist with uuid %s does not exist", uuid)
}

func GetAttachment(attachments []*types.Attachment, uuid string) (*types.Attachment, error) {
	for _, a := range attachments {
		if a.UUID == uuid {
			return a, nil
		}
	}
	return nil, fmt.Errorf("attachment with uuid %s does not exist", uuid)
}

func GetDate(dates []*types.Date, uuid string) (*types.Date, error) {
	for _, date := range dates {
		if date.UUID == uuid {
			return date, nil
		}
	}
	return nil, fmt.Errorf("date with uuid %s does not exist", uuid)
}

func GetWorkspace(workspaces []*workspace.Workspace, uuid string) (*workspace.Workspace, error) {
	for _, w := range workspaces {
		if w.UUID == uuid {
			return w, nil
		}
	}
	return nil, fmt.Errorf("workspace with uuid %s does not exist", uuid)
}
