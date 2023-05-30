package task

import "sharkedule/database/types"

func (t *Task) Convert() *types.Task {
	return &t.Task
}
