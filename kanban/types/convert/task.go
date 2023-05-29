package convert

import (
	"sharkedule/kanban/column/task"
	KTypes "sharkedule/kanban/types"
)

func ConvertTask(ITask KTypes.ITask) (*task.Task, error) {
	if t, ok := ITask.(*task.Task); ok {
		return t, nil
	}
	return nil, nil
}
