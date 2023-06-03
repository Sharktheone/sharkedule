package column

import (
	"errors"
	"sharkedule/kanban/board/column/task"
)

func (c *Column) GetTask(uuid string) (*task.Task, error) {
	for _, t := range c.Tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (c *Column) New(name string) (*task.Task, error) {
	t := task.NewTask(name)
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Tasks = append(c.Tasks, t)
	if err := c.Save(); err != nil {
		return nil, err
	}
	return t, nil
}
