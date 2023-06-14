package kanbandb

import (
	"fmt"
	types2 "github.com/Sharktheone/sharkedule/kanban/types"
)

func NewTask(tasks []*types2.Task, column *types2.Column, name string) *types2.Task {
	task := types2.NewTask(name)
	column.Tasks = append(column.Tasks, task.UUID)
	tasks = append(tasks, task)
	return task
}

func GetTask(tasks []*types2.Task, uuid string) (*types2.Task, error) {
	for _, t := range tasks {
		if t.UUID == uuid {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task with uuid %s does not exist", uuid)
}

func SaveTask(tasks []*types2.Task, task *types2.Task) error {
	for i, t := range tasks {
		if t.UUID == task.UUID {
			tasks[i] = task
			return nil
		}
	}
	return fmt.Errorf("task with uuid %s does not exist", task.UUID)
}

func SaveTasks(tasks []*types2.Task, tasksToSave []*types2.Task) {
	tasks = tasksToSave
}

func MoveTask(column, toColumn *types2.Column, uuid string, toIndex int) error {
	var colFound bool

	for index, t := range column.Tasks {
		if t == uuid {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			colFound = true
			break
		}
	}

	if !colFound {
		return fmt.Errorf("task %s not found in column %s", uuid, column)
	}

	if toIndex > len(toColumn.Tasks) {
		for i := range toColumn.Tasks {
			if i == toIndex {
				toColumn.Tasks = append(toColumn.Tasks[:i], append([]string{uuid}, toColumn.Tasks[i+1:]...)...)
			}
		}
	} else {
		toColumn.Tasks = append(toColumn.Tasks, uuid)
	}
	return nil
}

func DeleteTask(tasks []*types2.Task, uuid string) error {
	for index, t := range tasks {
		if t.UUID == uuid {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting task %s not found", uuid)
}

func AddTagToTask(tasks []*types2.Task, task, tag string) error { // TODO: This should not be an task array
	for _, t := range tasks {
		if t.UUID == task {
			t.Tags = append(t.Tags, tag)
			return nil
		}
	}
	return fmt.Errorf("error while adding tag %s to task %s not found", tag, task)
}

func RemoveTagOnTask(task *types2.Task, tag string) error {
	for index, t := range task.Tags {
		if t == tag {
			task.Tags = append(task.Tags[:index], task.Tags[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error while deleting tag %s not found on task %s", tag, task.UUID)
}

func DeleteTaskOnColumn(column *types2.Column, uuid string) error {
	for index, t := range column.Tasks {
		if t == uuid {
			column.Tasks = append(column.Tasks[:index], column.Tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("failed to delete task on column: task %v not found on column %v", uuid, column.UUID)
}

func RenameTask(task *types2.Task, name string) {
	task.Name = name
}
