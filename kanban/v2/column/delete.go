package column

import "github.com/Sharktheone/sharkedule/database/db"

func (c *Column) Delete() error {
	return db.DBV2.DeleteColumn(c.UUID)
}
