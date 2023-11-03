package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/environment"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func Get(workspace, column string) (*Column, error) {
	c, err := db.DB.GetColumn(workspace, column)
	if err != nil {
		return nil, err
	}
	return &Column{
		Column:    c,
		Workspace: workspace,
	}, nil
}

func (c *Column) Env() *types.Environment {
	return environment.GetColumnEnv(&c.UUID)
}
