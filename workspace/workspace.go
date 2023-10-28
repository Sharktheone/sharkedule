package workspace

import "github.com/google/uuid"

type Workspace struct {
	UUID     string
	Name     string
	Tags     []string
	Projects []string

	//All uuids here as list?
}

func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Name: name,
		UUID: uuid.New().String(),
	}
}

//TODO implement all other functions like on user
