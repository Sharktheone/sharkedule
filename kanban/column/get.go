package column

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"github.com/Sharktheone/sharkedule/kanban/environment"
	"github.com/Sharktheone/sharkedule/kanban/types"
)

func Get(column string) (*Column, error) {
	c, err := db.DB.GetColumn(column)
	if err != nil {
		return nil, err
	}
	return &Column{
		Column: c,
	}, nil
}

func (c *Column) Env() *types.Environment {
	return environment.GetColumnEnv(&c.UUID)
}
