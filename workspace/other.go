package workspace

import (
	"github.com/Sharktheone/sharkedule/database/db"
	ktypes "github.com/Sharktheone/sharkedule/kanban/types"
)

//Other functions

func (w *Workspace) GetStatus(uuid string) (*ktypes.Status, error) {
	return db.DB.GetStatus(w.UUID, uuid)
}

func (w *Workspace) GetPriority(uuid string) (*ktypes.Priority, error) {
	return db.DB.GetPriority(w.UUID, uuid)
}

func (w *Workspace) GetChecklist(uuid string) (*ktypes.Checklist, error) {
	return db.DB.GetChecklist(w.UUID, uuid)
}

func (w *Workspace) GetAttachment(uuid string) (*ktypes.Attachment, error) {
	return db.DB.GetAttachment(w.UUID, uuid)
}

func (w *Workspace) GetDate(uuid string) (*ktypes.Date, error) {
	return db.DB.GetDate(w.UUID, uuid)
}

//GetUser(uuid string) (types.Member, error) TODO
