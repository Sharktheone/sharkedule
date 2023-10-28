package project

import "github.com/google/uuid"

type Project struct {
	UUID    string
	Name    string
	Tags    []string
	Tasks   []string
	Columns []string
	Boards  []string
}

func NewProject(name string) *Project {
	return &Project{
		Name: name,
		UUID: uuid.New().String(),
	}
}

//TODO implement all other functions like on user
