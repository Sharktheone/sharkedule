package task

import "github.com/Sharktheone/sharkedule/database/db"

func Get(uuid string) (*Task, error) {
	t, err := db.DBV2.GetTask(uuid)
	if err != nil {
		return nil, err
	}

	return &Task{
		Task: t,
	}, nil
}
