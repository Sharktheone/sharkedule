package element

import "github.com/Sharktheone/sharkedule/workspace"

type Reference struct {
	UUID   string   `json:"uuid" bson:"uuid" yaml:"uuid"`
	Field  []string `json:"field" bson:"field" yaml:"field"`
	Linked bool     `json:"linked" bson:"linked" yaml:"linked"`
}

type ReferenceGroup struct {
	refs []Reference
	ws   *workspace.Workspace //Workspace Handle
	e    *Element             //Element Handle
}

func (rg *ReferenceGroup) AddReference(r Reference) {
	rg.ws.GetElement(r.UUID).ReferencedBy.AddReference(Reference{
		UUID:   rg.e.UUID,
		Field:  r.Field,
		Linked: r.Linked,
	})

	rg.refs = append(rg.refs, r)
}
