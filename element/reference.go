package element

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

type Reference struct {
	UUID   string   `json:"uuid" bson:"uuid" yaml:"uuid"`
	Field  []string `json:"field" bson:"field" yaml:"field"`
	Linked bool     `json:"linked" bson:"linked" yaml:"linked"`
}

func (r Reference) GetUUID() string {
	return r.UUID
}

func (r Reference) GetField() []string {
	return r.Field
}

func (r Reference) GetLinked() bool {
	return r.Linked
}

type ReferenceGroup struct {
	refs []Reference
	ws   string //Workspace Handle
	e    string //Element Handle
}

func (rg *ReferenceGroup) AddReference(r Reference) {
	elem, err := db.DB.GetElement(rg.ws, r.UUID)
	if err != nil {
		return
	}

	elem.GetReferencedBy().AddReference(Reference{
		UUID:   rg.e,
		Field:  r.Field,
		Linked: r.Linked,
	})

	rg.refs = append(rg.refs, r)
}
