package project

import (
	"github.com/Sharktheone/sharkedule/element/elements"
	"github.com/Sharktheone/sharkedule/field"
	"github.com/google/uuid"
)

type Project struct {
	UUID    string
	Name    string
	Tags    []string
	Tasks   []string
	Columns []string
	Boards  []string

	Fields []field.Field
}

func NewProject(name string) *Project {
	return &Project{
		Name: name,
		UUID: uuid.New().String(),
	}
}

//TODO implement all other functions like on user

func (p *Project) GetUUID() string {
	return p.UUID
}

func (p *Project) GetName() string {
	return p.Name
}

func (p *Project) GetFields() []field.Field {
	return p.Fields
}

func (p *Project) GetType() elements.Type {
	return elements.Project
}
