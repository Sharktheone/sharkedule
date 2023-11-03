package column

import (
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/google/uuid"
)

func NewColumn(workspace, name string) *Column {
	return &Column{
		Column: &types.Column{
			UUID: uuid.New().String(),
			Name: name,
		},
		Workspace: workspace,
	}
}
