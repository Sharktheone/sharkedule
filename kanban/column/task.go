package column

import (
	"errors"
	"sharkedule/kanban/column/task"
)

func (c *Column) GetTask(uuid string) (*task.Task, int, error) {
	for index, t := range c.Tasks {
		if t.UUID == uuid {
			return t, index, nil
		}
	}
	return nil, -1, errors.New("task not found")
}
