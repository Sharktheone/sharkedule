package column

import "github.com/Sharktheone/sharkedule/database/db"

func (c *Column) Save() error {
	return db.DB.SaveColumn(c.Column)
}
