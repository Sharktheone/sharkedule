package element

import (
	"github.com/Sharktheone/sharkedule/database/db"
)

type Reference struct {
	UUID   string   `json:"uuid" bson:"uuid" yaml:"uuid"`
	Field  []string `json:"field" bson:"field" yaml:"field"`
	Linked bool     `json:"linked" bson:"linked" yaml:"linked"`
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

	elem.ReferencedBy.AddReference(Reference{
		UUID:   rg.e,
		Field:  r.Field,
		Linked: r.Linked,
	})

	rg.refs = append(rg.refs, r)
}
