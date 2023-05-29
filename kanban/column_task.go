package kanban

import (
	"errors"
)

func (c *Column) GetTask(uuid string) (*Task, error) {
	for _, t := range c.Tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (c *Column) New(name string) *Task {
	t := NewTask(name)
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Tasks = append(c.Tasks, t)
	return t
}
