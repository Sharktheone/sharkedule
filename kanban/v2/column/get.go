package column

import "github.com/Sharktheone/sharkedule/database/db"

func Get(column string) (*Column, error) {
	c, err := db.DBV2.GetColumn(column)
	if err != nil {
		return nil, err
	}
	return &Column{
		Column: c,
	}, nil
}
