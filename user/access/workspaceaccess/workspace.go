package workspaceaccess

import (
	"errors"
	"github.com/Sharktheone/sharkedule/user/permissions"
)

type WorkspaceAccess struct {
	UUID        string                     `json:"uuid"`
	Permissions permissions.WorkspacePerms `json:"permissions"`

	//TODO: roles / groups

	AllBoards bool          `json:"all_boards"`
	Boards    []BoardAccess `json:"boards"`

	AllColumns bool           `json:"all_columns"`
	Columns    []ColumnAccess `json:"columns"`

	AllTasks bool         `json:"all_tasks"`
	Tasks    []TaskAccess `json:"tasks"`

	AllTags bool        `json:"all_tags"`
	Tags    []TagAccess `json:"tags"`

	AllOrgs bool        `json:"all_orgs"`
	Orgs    []OrgAccess `json:"orgs"`

	AllStatuses bool           `json:"all_statuses"`
	Statuses    []StatusAccess `json:"statuses"`

	AllPriorities bool             `json:"all_priorities"`
	Priorities    []PriorityAccess `json:"priorities"`

	AllChecklists bool              `json:"all_checklists"`
	Checklists    []ChecklistAccess `json:"checklists"`

	AllAttachments bool               `json:"all_attachments"`
	Attachments    []AttachmentAccess `json:"attachments"`

	AllDates bool         `json:"all_dates"`
	Dates    []DateAccess `json:"dates"`
}

type BoardAccess struct {
	UUID        string
	AllColumns  bool
	Columns     []ColumnAccess
	Permissions permissions.BoardPerms //TODO: Override permissions for specific boards
}

type ColumnAccess struct {
	UUID        string
	AllTasks    bool
	Tasks       []TaskAccess
	Permissions permissions.ColumnPerms //TODO: Override permissions for specific columns
}

type TaskAccess struct {
	UUID        string
	Permissions permissions.TaskPerms //TODO: Override permissions for specific tasks
}

type TagAccess struct {
	UUID        string
	Permissions permissions.TagPerms //TODO: Override permissions for specific tags
}

type OrgAccess struct {
	UUID        string
	Permissions permissions.OrgPerms
}

type CreateAccess struct {
	Boards bool
	Orgs   bool
}

type StatusAccess struct {
	UUID        string
	Permissions permissions.StatusPerms
}

type PriorityAccess struct {
	UUID        string
	Permissions permissions.PriorityPerms
}

type ChecklistAccess struct {
	UUID        string
	Permissions permissions.ChecklistPerms
}

type AttachmentAccess struct {
	UUID        string
	Permissions permissions.AttachmentPerms
}

type DateAccess struct {
	UUID        string
	Permissions permissions.DatePerms
}

func (wa *WorkspaceAccess) Board(uuid string) (*BoardAccess, error) {

	for _, b := range wa.Boards {
		if b.UUID == uuid {
			return &b, nil
		}
	}
	return nil, errors.New("board not found")
}

func (wa *WorkspaceAccess) Column(uuid string) (*ColumnAccess, error) {
	for _, c := range wa.Columns {
		if c.UUID == uuid {
			return &c, nil
		}
	}
	return nil, errors.New("column not found")
}

func (wa *WorkspaceAccess) Task(uuid string) (*TaskAccess, error) {
	for _, t := range wa.Tasks {
		if t.UUID == uuid {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (wa *WorkspaceAccess) Tag(uuid string) (*TagAccess, error) {
	for _, t := range wa.Tags {
		if t.UUID == uuid {
			return &t, nil
		}
	}
	return nil, errors.New("tag not found")
}

func (wa *WorkspaceAccess) Status(uuid string) (*StatusAccess, error) {
	for _, s := range wa.Statuses {
		if s.UUID == uuid {
			return &s, nil
		}
	}
	return nil, errors.New("status not found")
}

func (wa *WorkspaceAccess) Priority(uuid string) (*PriorityAccess, error) {
	for _, p := range wa.Priorities {
		if p.UUID == uuid {
			return &p, nil
		}
	}
	return nil, errors.New("priority not found")
}

func (wa *WorkspaceAccess) Checklist(uuid string) (*ChecklistAccess, error) {
	for _, c := range wa.Checklists {
		if c.UUID == uuid {
			return &c, nil
		}
	}
	return nil, errors.New("checklist not found")
}

func (wa *WorkspaceAccess) Attachment(uuid string) (*AttachmentAccess, error) {
	for _, a := range wa.Attachments {
		if a.UUID == uuid {
			return &a, nil
		}
	}
	return nil, errors.New("attachment not found")
}

func (wa *WorkspaceAccess) Date(uuid string) (*DateAccess, error) {
	for _, d := range wa.Dates {
		if d.UUID == uuid {
			return &d, nil
		}
	}
	return nil, errors.New("date not found")
}
