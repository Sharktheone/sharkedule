package column

import "C"
import "github.com/Sharktheone/sharkedule/database/db"

func (c *Column) Rename(name string) error {
	return db.DB.RenameColumn(name)
}
