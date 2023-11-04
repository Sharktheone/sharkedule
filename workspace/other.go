package workspace

import "github.com/Sharktheone/sharkedule/kanban/types"

//Other functions

func (w *Workspace) GetStatus(workspace, uuid string) (*types.Status, error) {

}

func (w *Workspace) GetPriority(workspace, uuid string) (*types.Priority, error) {

}

func (w *Workspace) GetChecklist(workspace, uuid string) (*types.Checklist, error) {

}

func (w *Workspace) GetAttachment(workspace, uuid string) (*types.Attachment, error) {

}

func (w *Workspace) GetDate(workspace, uuid string) (*types.Date, error) {

}

//GetUser(uuid string) (*types.Member, error) TODO
