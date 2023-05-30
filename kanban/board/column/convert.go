package column

import "sharkedule/database/types"

func (c *Column) Convert() *types.Column {
	return c.Column
}
