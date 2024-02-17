package workspace

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/types"
)

func (w *Workspace) GetElement(u string) (*types.Element, error) {
	return db.DB.GetElement(w.UUID, u)
}

func (w *Workspace) CreateElement(e *types.ElementType, name string) (*types.Element, error) {
	return db.DB.CreateElement(w.UUID, e, name)
}

func (w *Workspace) GetElements(u []string) ([]*types.Element, error) {
	return db.DB.GetElements(w.UUID, u)
}
